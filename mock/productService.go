package mock

import "github.com/e2dk4r/supermarket"

type ProductService struct {
	ProductFn      func(id string) (*supermarket.Product, error)
	productInvoked bool

	ProductsFn      func(page int, perPage int) ([]*supermarket.Product, error)
	productsInvoked bool

	CreateProductFn      func(p *supermarket.Product) error
	createProductInvoked bool

	DeleteProductFn      func(id string) (bool, error)
	deleteProductInvoked bool

	IsDuplicateErrorFn      func(err error) bool
	isDuplicateErrorInvoked bool
}

func (ps *ProductService) Product(id string) (*supermarket.Product, error) {
	ps.productInvoked = true
	return ps.ProductFn(id)
}

func (ps *ProductService) ProductInvoked() bool {
	return ps.productInvoked
}

func (ps *ProductService) Products(page int, perPage int) ([]*supermarket.Product, error) {
	ps.productsInvoked = true
	return ps.ProductsFn(page, perPage)
}

func (ps *ProductService) ProductsInvoked() bool {
	return ps.productsInvoked
}

func (ps *ProductService) CreateProduct(p *supermarket.Product) error {
	ps.createProductInvoked = true
	return ps.CreateProductFn(p)
}

func (ps *ProductService) CreateProductInvoked() bool {
	return ps.createProductInvoked
}

func (ps *ProductService) DeleteProduct(id string) (bool, error) {
	ps.deleteProductInvoked = true
	return ps.DeleteProductFn(id)
}

func (ps *ProductService) DeleteProductInvoked() bool {
	return ps.deleteProductInvoked
}

func (ps *ProductService) IsDuplicateError(err error) bool {
	ps.isDuplicateErrorInvoked = true
	return ps.IsDuplicateErrorFn(err)
}

func (ps *ProductService) IsDuplicateErrorInvoked() bool {
	return ps.isDuplicateErrorInvoked
}
