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
	// it doesn't create product ID but in-memory db doesn't need to
	db.Products = append(db.Products, p)
}

func (db *InMemoryDB) EditProduct(p model.Product) error {
	for i, product := range db.Products {
		if product.ID == p.ID {
			db.Products[i] = p
			return nil
		}
	}
	return cannotFindProductError(p.ID)
}

func (db *InMemoryDB) DeleteProduct(p model.Product) error {
	for i, product := range db.Products {
		if product.ID == p.ID {
			db.Products = append(db.Products[:i], db.Products[i+1:]...)
			return nil
		}
	}
	return cannotFindProductError(p.ID)
}

func cannotFindProductError(id string) error {
	return errors.New(fmt.Sprintf("Cannot find product with ID %s in database.", id))
}
