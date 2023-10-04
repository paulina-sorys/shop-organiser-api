// Package server handles all api endpoints for managing grocery and chemical products in the application.
// It follows CRUD notion. As it's a work-in-progress package it uses in-memory database for simplicity.
// It's intended to switch into database such as Postgress/MongoDB or other later on.
package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/paulina-sorys/shop-organiser-api/model"
)

// InMemoryDB interface holds all operations you can impose on database
type InMemoryDB interface {
	GetAllProducts() []model.Product
	AddProduct(model.Product)
}

type Server struct {
	http.Handler // handler for all api endpoints
	store        InMemoryDB
}

func (s *Server) allProductsHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(s.store.GetAllProducts())
	w.WriteHeader(http.StatusOK)
}

func (s *Server) newProductHandler(w http.ResponseWriter, req *http.Request) {
	var product model.Product
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		log.Fatalf("Unable to parse request body %q into Product object, '%v'", req.Body, err)
	}
	s.store.AddProduct(product)
	w.WriteHeader(http.StatusAccepted)
}

// New creates all api endpoints handlers. It requires an instance of database.
func New(db InMemoryDB) *Server {
	s := Server{}
	s.store = db

	mux := http.NewServeMux()
	mux.Handle("/api/v1/products/all", http.HandlerFunc(s.allProductsHandler))
	mux.Handle("/api/v1/product/new", http.HandlerFunc(s.newProductHandler))

	s.Handler = mux
	return &s
}
