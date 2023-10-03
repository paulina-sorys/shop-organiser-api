package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/paulina-sorys/shop-organiser-api/model"
	"github.com/stretchr/testify/assert"
)

type StubInMemeoryDB struct {
	products []model.Product
}

func (db *StubInMemeoryDB) GetAllProducts() []model.Product {
	return db.products
}

func (db *StubInMemeoryDB) AddProduct(p model.Product) {
	db.products = append(db.products, p)
}

// Test fails if received request response status is not expected
func checkStatus(t *testing.T, want, got int) {
	assert.Equal(t, want, got, "Wrong status. Expected %d, got %d", want, got)
}

// Test fails if received slice of products is not expected
func checkProducts(t *testing.T, want, got []model.Product) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Received incorrect slice of products. Expected %q, got %q", want, got)
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
	db := &StubInMemeoryDB{
		[]model.Product{
			{Name: "juice"},
			{Name: "cheese"},
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
		productsBeforePOST := db.products

		productJSON, _ := json.Marshal(model.Product{Name: "chocolate"})
		response := server.callApi(http.MethodPost, "/api/v1/product/new", productJSON)

		checkStatus(t, http.StatusAccepted, response.Code)

		want := append(productsBeforePOST, model.Product{Name: "chocolate"})
		checkProducts(t, want, server.store.GetAllProducts())
	})
}
