package cockroachdb

import (
	"context"
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
	return nil, fmt.Errorf("no")
}

func (ps *ProductService) CreateProduct(p *supermarket.Product) error {
	return fmt.Errorf("no")
}

func (ps *ProductService) DeleteProduct(p *supermarket.Product) error {
	return fmt.Errorf("no")
}
