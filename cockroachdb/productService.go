package cockroachdb

import (
	"fmt"

	"github.com/e2dk4r/supermarket"

	"github.com/jackc/pgx/v4"
)

type ProductService struct {
	Conn *pgx.Conn
}

func (ps *ProductService) Product(id string) (*supermarket.Product, error) {
	return nil, fmt.Errorf("no")
}

func (ps *ProductService) Products() ([]*supermarket.Product, error) {
	return nil, fmt.Errorf("no")
}

func (ps *ProductService) CreateProduct(p *supermarket.Product) error {
	return fmt.Errorf("no")
}

func (ps *ProductService) DeleteProduct(p *supermarket.Product) error {
	return fmt.Errorf("no")
}
