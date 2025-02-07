package db

import "model"

type InMemoryDB struct {
	Products []model.Product
}

func (db *InMemoryDB) GetAllProducts() []model.Product {
	return db.Products
}

func (db *InMemoryDB) AddProduct(p model.Product) {
	db.Products = append(db.Products, p)
}

func (db *InMemoryDB) EditProduct(p model.Product) {
	for i, product := range db.Products {
		if product.ID == p.ID {
			db.Products[i] = p
			break
		}
	}
}
