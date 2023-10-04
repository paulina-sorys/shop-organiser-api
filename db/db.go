package db

import "github.com/paulina-sorys/shop-organiser-api/model"

type InMemoryDB struct {
	Products []model.Product
}

func (db *InMemoryDB) GetAllProducts() []model.Product {
	return db.Products
}

func (db *InMemoryDB) AddProduct(p model.Product) {
	db.Products = append(db.Products, p)
}
