package server

import (
	"bytes"
	"db"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"model"
)

// Test fails if received request response status is not expected
func checkStatus(t *testing.T, want, got int) {
	assert.Equal(t, want, got, "Wrong http status code.\nExpected: %d\n\t Got: %d", want, got)
}

// Test fails if received slice of products is not expected
func checkProducts(t *testing.T, want, got []model.Product) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Received incorrect slice of products.\nExpected: %q\n\t Got: %q", want, got)
	}
}

// Calls server endpoint and returns its response
func (s *Server) callApi(method, path string, body []byte) *httptest.ResponseRecorder {
	request, err := http.NewRequest(method, path, bytes.NewReader(body))
	if err != nil {
		log.Fatalf("Cannot create %s request to endpoint %s with body %q, '%v'", method, path, body, err)
	}
	response := httptest.NewRecorder()

	s.ServeHTTP(response, request)
	return response
}

func TestServer(t *testing.T) {
	db := &db.InMemoryDB{
		Products: []model.Product{
			{Name: "juice", ID: "123"},
			{Name: "cheese", ID: "342"},
		},
	}
	server := New(db)

	t.Run("GET all products", func(t *testing.T) {
		response := server.callApi(http.MethodGet, "/api/v1/products/all", nil)

		checkStatus(t, http.StatusOK, response.Code)

		var got []model.Product

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of products, '%v'", response.Body, err)
		}

		checkProducts(t, server.store.GetAllProducts(), got)
	})

	t.Run("POST new product", func(t *testing.T) {
		productsBeforePOST := db.Products
		productToAdd := model.Product{Name: "chocolate"}

		productJSON, _ := json.Marshal(productToAdd)
		response := server.callApi(http.MethodPost, "/api/v1/product/new", productJSON)

		checkStatus(t, http.StatusAccepted, response.Code)

		want := append(productsBeforePOST, productToAdd)
		checkProducts(t, want, server.store.GetAllProducts())
	})

	t.Run("try POST new product with unrecognised json representation of product", func(t *testing.T) {
		response := server.callApi(http.MethodPost, "/api/v1/product/new", []byte("Wrong product JSON representation"))
		checkStatus(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PUT product (edit existing product)", func(t *testing.T) {
		productEditOutcome := model.Product{Name: "orange juice", ID: "123"}
		productsAfterPUT := func() []model.Product {
			products := make([]model.Product, len(db.Products))
			copy(products, db.Products)
			products[0] = productEditOutcome
			return products
		}()
		productJSON, _ := json.Marshal(productEditOutcome)

		response := server.callApi(http.MethodPut, "/api/v1/product/edit", productJSON)

		checkStatus(t, http.StatusOK, response.Code)
		checkProducts(t, productsAfterPUT, server.store.GetAllProducts())
	})
}
