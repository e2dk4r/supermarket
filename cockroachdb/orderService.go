package cockroachdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/e2dk4r/supermarket"

	"github.com/jackc/pgx/v4"
)

type OrderService struct {
	Conn *pgx.Conn
}

func (s *OrderService) Order(id string) (*supermarket.Order, error) {
	rows, err := s.Conn.Query(context.Background(), `
	SELECT
		o.status,
		p.id,
		p.name,
		op.amount,
		p.price
	FROM orders o
	JOIN order_product op ON op.order_id   = o.id
	JOIN products p       ON op.product_id = p.id
	WHERE o.id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list := []*supermarket.OrderItem{}
	var status supermarket.OrderStatus

	for rows.Next() {
		var productId, name string
		var price float32
		var amount int

		err = rows.Scan(&status, &productId, &name, &amount, &price)
		if err != nil {
			return nil, err
		}

		list = append(list, &supermarket.OrderItem{
			Product: supermarket.Product{
				Id:    productId,
				Name:  name,
				Price: price,
			},
			Amount: amount,
		})
	}

	return &supermarket.Order{
		Id:     id,
		Status: supermarket.OrderStatus(status),
		Basket: list,
	}, nil
}

func (s *OrderService) Orders(page int, perPage int) ([]*supermarket.Order, error) {
	if page <= 0 || page > 500000 {
		return nil, errors.New("page limits exceeded")
	}
	if perPage <= 0 || perPage > 250 {
		return nil, errors.New("page limits exceeded")
	}

	offset := (page - 1) * perPage
	rows, err := s.Conn.Query(context.Background(), `
	SELECT
		o.id,
		o.status,
		p.id,
		op.amount,
		p.price
	FROM orders o
	JOIN order_product op ON o.id = op.order_id
	JOIN products p       ON op.product_id = p.id
	OFFSET $1 LIMIT $2
	`, offset, perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := map[string]*supermarket.Order{}

	for rows.Next() {
		var id, productId string
		var price float32
		var status, amount int

		err = rows.Scan(&id, &status, &productId, &amount, &price)
		if err != nil {
			return nil, err
		}

		order, ok := list[id]
		if ok {
			order.Basket = append(order.Basket, &supermarket.OrderItem{
				Product: supermarket.Product{
					Id:    productId,
					Price: price,
				},
				Amount: amount,
			})
			continue
		}
		order = &supermarket.Order{
			Id:     id,
			Status: supermarket.OrderStatus(status),
			Basket: []*supermarket.OrderItem{
				{
					Product: supermarket.Product{
						Id:    productId,
						Price: price,
					},
					Amount: amount,
				},
			},
		}
		list[id] = order
	}

	if len(list) == 0 {
		return []*supermarket.Order{}, nil
	}

	orders := make([]*supermarket.Order, 0, len(list))
	for _, order := range list {
		order.Total = 0.0
		for _, product := range order.Basket {
			order.Total += product.Product.Price * float32(product.Amount)
		}

		orders = append(orders, order)
	}

	return orders, nil
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
