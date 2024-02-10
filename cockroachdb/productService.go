package cockroachdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/e2dk4r/supermarket"

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
	return fmt.Errorf("no")
}

func (ps *ProductService) DeleteProduct(p *supermarket.Product) error {
	return fmt.Errorf("no")
}
