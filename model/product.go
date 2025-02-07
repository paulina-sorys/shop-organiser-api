package model

type Product struct {
	Name string
	ID   string
}

// Store interface holds all operations you can impose on database
type Store interface {
	GetAllProducts() []Product
	AddProduct(Product)
	EditProduct(Product)
}
