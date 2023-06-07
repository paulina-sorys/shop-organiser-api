package db

import "github.com/paulina-sorys/shop-organiser/model"

type InMemeoryDB struct {
	Products []model.Product
}

func (db *InMemeoryDB) GetAllProducts() []model.Product {
	return db.Products
}

func (db *InMemeoryDB) AddProduct(p model.Product) {
	db.Products = append(db.Products, p)
}
