package server

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Name string
}

func New(w http.ResponseWriter, r *http.Request) {
	allProducts := []Product{{"juice"}, {"cheese"}}
	json.NewEncoder(w).Encode(allProducts)
}
