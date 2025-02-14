package model

type Product struct {
	Name string
	ID   int
}

// Store interface holds all operations you can impose on database
type Store interface {
	GetAllProducts() []Product
	AddProduct(Product)
	EditProduct(Product) error
	DeleteProduct(Product) error
}
