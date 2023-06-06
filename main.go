package main

import (
	"log"
	"net/http"

	"github.com/paulina-sorys/shop-organiser/server"
)

type InMemeoryDB struct {
	products []server.Product
}

func (db *InMemeoryDB) GetAllProducts() []server.Product {
	return db.products
}

func (db *InMemeoryDB) AddProduct(p server.Product) {
	db.products = append(db.products, p)
}

func main() {
	db := &InMemeoryDB{[]server.Product{{Name: "test"}}}
	log.Fatal(http.ListenAndServe(":5000", server.New(db).Handler))
}
