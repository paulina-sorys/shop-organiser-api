package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test fails if received request response status is not expected
func checkStatus(t *testing.T, expected, got int) {
	assert.Equal(t, expected, got, "Wrong status. Expected %d, got %d", expected, got)
}

// Calls server endpoint and return its response
func (s *Server) callApi(method, path string, body []byte) *httptest.ResponseRecorder {
	request, err := http.NewRequest(method, path, bytes.NewReader(body))
	if err != nil {
		log.Fatalf("Cannot create %s request to endpoint %s with body %q, '%v'", method, path, body, err)
	}
	response := httptest.NewRecorder()

	s.ServeHTTP(response, request)
	return response
}

type StubInMemeoryDB struct {
	products []Product
}

func (db *StubInMemeoryDB) GetAllProducts() []Product {
	return db.products
}

func (db *StubInMemeoryDB) AddProduct(p Product) {
	db.products = append(db.products, p)
}

func TestServer(t *testing.T) {
	db := &StubInMemeoryDB{
		[]Product{
			{"juice"},
			{"cheese"},
		},
	}
	server := New(db)

	t.Run("GET all products", func(t *testing.T) {
		response := server.callApi(http.MethodGet, "/api/v1/products/all", nil)

		checkStatus(t, http.StatusOK, response.Code)

		var got []Product

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of products, '%v'", response.Body, err)
		}

		want := server.store.GetAllProducts()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("POST new product", func(t *testing.T) {
		productsBeforePOST := server.store.GetAllProducts()

		productJSON, _ := json.Marshal(Product{"chocolate"})

		response := server.callApi(http.MethodPost, "/api/v1/product/new", productJSON)

		want := append(productsBeforePOST, Product{"chocolate"})

		if !reflect.DeepEqual(want, server.store.GetAllProducts()) {
			t.Errorf("got %q, want %q", server.store.GetAllProducts(), want)
		}

		checkStatus(t, http.StatusAccepted, response.Code)
	})

}
