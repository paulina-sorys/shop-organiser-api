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
	lastProductID := db.Products[len(db.Products)-1].ID
	p.ID = lastProductID + 1
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

func cannotFindProductError(id int) error {
	return errors.New(fmt.Sprintf("Cannot find product with ID %d in database.", id))
}
