package supermarket

type Product struct {
	Id    string
	Name  string
	Price float32
}

type ProductService interface {
	Product(id string) (*Product, error)
	Products(page int, perPage int) ([]*Product, error)
	CreateProduct(p *Product) error
	DeleteProduct(p *Product) error
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
