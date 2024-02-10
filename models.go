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
	DeleteProduct(id string) (bool, error)
	IsDuplicateError(err error) bool
	IsNotFoundError(err error) bool
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

type User struct {
	Id       string `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type UserService interface {
	User(username string) (*User, error)

	// CreateUser creates user using username and password
	// field of model. And writes generated id to id field
	CreateUser(user *User) error
}

type AuthService interface {
	// CreateToken generates an authentication token.
	// It returns authentication token if user exists and
	// password is a match.
	CreateToken(u *User) (string, error)

	// VerifyToken verifies an authentication token specified.
	// It returns true if token is correct, false otherwise
	VerifyToken(key string) error

	// CreateAnonToken generates an short lived authentication token.
	CreateAnonToken() (string, error)

	// VerifyAnonToken verifies an authentication token specified.
	// It returns nil if token is correct
	VerifyAnonToken(key string) error
}

type PasswordService interface {
	// Hash generates hashed password from plain-text password
	Hash(password string) (string, error)

	// Verify performs compares between a plain-text password and hash,
	// using the parameters and salt contained in the hash. It returns true if
	// they match, otherwise it returns false.
	Verify(password string, hash string) (bool, error)
}

type RandomService interface {
	GenerateString(n int) (string, error)
}

type DbErrorService interface {
	IsDuplicateError(err error) bool
	IsNotFoundError(err error) bool
}
