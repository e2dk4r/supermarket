package cockroachdb

import (
	"fmt"

	"github.com/e2dk4r/supermarket"

	"github.com/jackc/pgx/v4"
)

type OrderService struct {
	Conn *pgx.Conn
}

func (s *OrderService) Order(id string) (*supermarket.Order, error) {
	return nil, fmt.Errorf("no")
}

func (s *OrderService) Orders() ([]*supermarket.Order, error) {
	return nil, fmt.Errorf("no")
}

func (s *OrderService) CreateOrder(o *supermarket.Order) error {
	return fmt.Errorf("no")
}

func (s *OrderService) DeleteOrder(o *supermarket.Order) error {
	return fmt.Errorf("no")
}

func (s *OrderService) OrderBasket(o *supermarket.Order) error {
	return fmt.Errorf("no")
}
