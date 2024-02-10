package supermarket

import "errors"

type Product struct {
	Id    string
	Name  string
	Price float32
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}
	if len(p.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}
	if len(p.Name) > 1000 {
		return errors.New("name must be at max 1000 characters")
	}

	if p.Price <= 0 {
		return errors.New("price is required and must be greater than zero")
	}

	return nil
}

type ProductService interface {
	Product(id string) (*Product, error)
	Products(page int, perPage int) ([]*Product, error)
	CreateProduct(p *Product) error
	DeleteProduct(p *Product) error
	IsDuplicateError(err error) bool
}

type Order struct {
	Id     string
	Basket []OrderItem
}

type OrderItem struct {
	Amount  int
	Product Product
}

type OrderService interface {
	Order(id string) (*Order, error)
	Orders() ([]*Order, error)
	CreateOrder(o *Order) error
	DeleteOrder(o *Order) error
	OrderBasket(o *Order) error
}
