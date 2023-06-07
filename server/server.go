package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/paulina-sorys/shop-organiser/model"
)

type InMemoryDB interface {
	GetAllProducts() []model.Product
	AddProduct(model.Product)
}

type Server struct {
	http.Handler
	store InMemoryDB
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

func New(db InMemoryDB) *Server {
	s := Server{}
	s.store = db

	mux := http.NewServeMux()
	mux.Handle("/api/v1/products/all", http.HandlerFunc(s.allProductsHandler))
	mux.Handle("/api/v1/product/new", http.HandlerFunc(s.newProductHandler))

	s.Handler = mux
	return &s
}
