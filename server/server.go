// Package server handles all api endpoints for managing grocery and chemical products in the application.
// It follows CRUD notion. As it's a work-in-progress package it uses in-memory database for simplicity.
// It's intended to switch into database such as Postgress/MongoDB or other later on.
package server

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
)

type Server struct {
	http.Handler // handler for all api endpoints
	store        model.Store
}

// New creates all api endpoints handlers.
func New(db model.Store) *Server {
	s := Server{}
	s.store = db

	mux := http.NewServeMux()
	mux.Handle("/api/v1/products/all", http.HandlerFunc(s.allProductsHandler))
	mux.Handle("/api/v1/product/new", http.HandlerFunc(s.newProductHandler))
	mux.Handle("/api/v1/product/edit", http.HandlerFunc(s.editProductHandler))

	s.Handler = mux
	return &s
}

func (s *Server) allProductsHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(s.store.GetAllProducts())
	w.WriteHeader(http.StatusOK)
}

func (s *Server) newProductHandler(w http.ResponseWriter, req *http.Request) {
	var product model.Product
	err := json.NewDecoder(req.Body).Decode(&product)
	handleDecodingProductError(err, w)
	s.store.AddProduct(product)
	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) editProductHandler(w http.ResponseWriter, req *http.Request) {
	var product model.Product
	err := json.NewDecoder(req.Body).Decode(&product)
	handleDecodingProductError(err, w)
	productEditError := s.store.EditProduct(product)
	handleEditProductError(productEditError, w)
	w.WriteHeader(http.StatusOK)
}

func handleDecodingProductError(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to parse request body into Product object, err: '%v'", err)
		return
	}
}

func handleEditProductError(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, err.Error())
		return
	}
}
