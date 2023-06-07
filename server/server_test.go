package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/paulina-sorys/shop-organiser/model"
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
func checkStatus(t *testing.T, expected, got int) {
	assert.Equal(t, expected, got, "Wrong status. Expected %d, got %d", expected, got)
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
			{"juice"},
			{"cheese"},
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

		want := server.store.GetAllProducts()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("POST new product", func(t *testing.T) {
		productsBeforePOST := server.store.GetAllProducts()

		productJSON, _ := json.Marshal(model.Product{Name: "chocolate"})

		response := server.callApi(http.MethodPost, "/api/v1/product/new", productJSON)

		want := append(productsBeforePOST, model.Product{Name: "chocolate"})

		if !reflect.DeepEqual(want, server.store.GetAllProducts()) {
			t.Errorf("got %q, want %q", server.store.GetAllProducts(), want)
		}

		checkStatus(t, http.StatusAccepted, response.Code)
	})

}
