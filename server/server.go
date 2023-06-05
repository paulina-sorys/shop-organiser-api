package server

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name string
}

type Server struct {
	http.Handler
}

func (s *Server) allProductsHandler(w http.ResponseWriter, req *http.Request) {
	allProducts := []Product{{"juice"}, {"cheese"}}
	json.NewEncoder(w).Encode(allProducts)
}

func (s *Server) newProductHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func New() *Server {
	s := &Server{}

	mux := http.NewServeMux()
	mux.Handle("/api/v1/products/all", http.HandlerFunc(s.allProductsHandler))
	mux.Handle("/api/v1/product/new", http.HandlerFunc(s.newProductHandler))

	s.Handler = mux
	return s
}
