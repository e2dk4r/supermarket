package cockroachdb

import (
	"context"
	"errors"
	"log"

	"github.com/e2dk4r/supermarket"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type ProductService struct {
	Conn *pgx.Conn
}

func (ps *ProductService) Product(id string) (*supermarket.Product, error) {
	var name string
	var price float32

	row := ps.Conn.QueryRow(context.Background(), "SELECT name, price FROM products WHERE id = $1", id)
	err := row.Scan(&name, &price)
	if err != nil {
		return nil, err
	}

	return &supermarket.Product{
		Id:    id,
		Name:  name,
		Price: price,
	}, nil
}

func (ps *ProductService) Products(page int, perPage int) ([]*supermarket.Product, error) {
	if page <= 0 || page > 500000 {
		return nil, errors.New("page limits exceeded")
	}
	if perPage <= 0 || perPage > 250 {
		return nil, errors.New("page limits exceeded")
	}

	offset := (page - 1) * perPage
	rows, err := ps.Conn.Query(context.Background(), "SELECT id, name, price FROM products OFFSET $1 LIMIT $2", offset, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*supermarket.Product

	for rows.Next() {
		var id, name string
		var price float32

		err = rows.Scan(&id, &name, &price)
		if err != nil {
			return nil, err
		}

		list = append(list, &supermarket.Product{
			Id:    id,
			Name:  name,
			Price: price,
		})
	}

	return list, nil
}

func (ps *ProductService) CreateProduct(p *supermarket.Product) error {
	err := crdbpgx.ExecuteTx(context.Background(), ps.Conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		err := tx.QueryRow(context.Background(), "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id;", p.Name, p.Price).Scan(&p.Id)
		if err != nil {
			return err
		}

		log.Printf("product created: %#v", p)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) DeleteProduct(productId string) (bool, error) {
	deleted := false
	err := crdbpgx.ExecuteTx(context.Background(), ps.Conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		tag, err := tx.Exec(context.Background(), "DELETE FROM products WHERE id = $1", productId)
		if err != nil {
			return err
		}

		deleted = tag.RowsAffected() == 1
		return nil
	})

	if deleted {
		log.Printf("product deleted: %v", productId)
	}
	return deleted, err
}

func (ps *ProductService) IsDuplicateError(err error) bool {
	var pgError *pgconn.PgError

	if !errors.As(err, &pgError) {
		return false
	}

	return pgError.Code == "23505"
}
