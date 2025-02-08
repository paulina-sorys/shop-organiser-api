package db

import (
	"errors"
	"fmt"
	"model"
)

type InMemoryDB struct {
	Products []model.Product
}

func (db *InMemoryDB) GetAllProducts() []model.Product {
	return db.Products
}

func (db *InMemoryDB) AddProduct(p model.Product) {
	db.Products = append(db.Products, p)
}

func (db *InMemoryDB) EditProduct(p model.Product) error {
	for i, product := range db.Products {
		if product.ID == p.ID {
			db.Products[i] = p
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Cannot find product with ID %s in database.", p.ID))
}
