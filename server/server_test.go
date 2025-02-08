package server

import (
	"bytes"
	"db"
	"encoding/json"
	"fmt"
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
	assert.Equal(t, want, got, "Wrong http status code.\nExpected: %d\nGot\t: %d", want, got)
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
	dbStub := &db.InMemoryDB{
		Products: []model.Product{
			{Name: "juice", ID: 1},
			{Name: "cheese", ID: 2},
		},
	}
	server := New(dbStub)

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
		productsBeforePOST := dbStub.Products
		productToAdd := model.Product{Name: "chocolate"}

		productJSON, _ := json.Marshal(productToAdd)
		response := server.callApi(http.MethodPost, "/api/v1/product/new", productJSON)

		checkStatus(t, http.StatusOK, response.Code)

		productToAdd.ID = 3
		want := append(productsBeforePOST, productToAdd)
		checkProducts(t, want, server.store.GetAllProducts())
	})

	t.Run("try POST new product with unrecognised json representation of product", func(t *testing.T) {
		response := server.callApi(http.MethodPost, "/api/v1/product/new", []byte("Wrong product JSON representation"))
		checkStatus(t, http.StatusBadRequest, response.Code)
	})

	t.Run("PUT product (edit existing product)", func(t *testing.T) {
		productEditOutcome := model.Product{Name: "orange juice", ID: 1}
		productsAfterPUT := func() []model.Product {
			products := make([]model.Product, len(dbStub.Products))
			copy(products, dbStub.Products)
			products[0] = productEditOutcome
			return products
		}()
		productJSON, _ := json.Marshal(productEditOutcome)

		response := server.callApi(http.MethodPut, "/api/v1/product/edit", productJSON)

		checkStatus(t, http.StatusOK, response.Code)
		checkProducts(t, productsAfterPUT, server.store.GetAllProducts())
	})

	t.Run("try PUT product but product is not found in database", func(t *testing.T) {
		productNotInDatabase := model.Product{Name: "milk", ID: 456} // TODO: should validate for existence of ID before calling db
		productJSON, _ := json.Marshal(productNotInDatabase)
		response := server.callApi(http.MethodPut, "/api/v1/product/edit", productJSON)
		checkStatus(t, http.StatusUnprocessableEntity, response.Code)
	})

	t.Run("DELETE product", func(t *testing.T) {
		productToDelete := model.Product{Name: "orange juice", ID: 1} // TODO: should validate for existence of ID before calling db
		productsAfterDelete := func() []model.Product {
			products := make([]model.Product, len(dbStub.Products)-1)
			copy(products, dbStub.Products[1:])
			return products
		}()
		productJSON, _ := json.Marshal(productToDelete)
		response := server.callApi(http.MethodDelete, "/api/v1/product/delete", productJSON)
		checkStatus(t, http.StatusOK, response.Code)
		checkProducts(t, productsAfterDelete, server.store.GetAllProducts())
	})

	t.Run("DELETE not existing product", func(t *testing.T) {
		productNotInDatabase := model.Product{Name: "pencils", ID: 789}
		productJSON, _ := json.Marshal(productNotInDatabase)
		response := server.callApi(http.MethodDelete, "/api/v1/product/delete", productJSON)
		checkStatus(t, http.StatusUnprocessableEntity, response.Code)
	})
}

func TestCallingServerEndpointsWithIncorrectHttpMethods(t *testing.T) {
	productToAddJSON, _ := json.Marshal(model.Product{Name: "plant"})
	productEditOutcomeJSON, _ := json.Marshal(model.Product{Name: "juice", ID: 1})
	productToDeleteJSON, _ := json.Marshal(model.Product{Name: "cheese", ID: 2})

	dbStub := &db.InMemoryDB{
		Products: []model.Product{
			{Name: "juice", ID: 1},
			{Name: "cheese", ID: 2},
		},
	}
	server := New(dbStub)

	testCases := []struct {
		description string
		apiPath     string
		json        []byte
	}{
		{"getting all products", "/api/v1/products/all", nil},
		{"adding new product", "/api/v1/product/new", productToAddJSON},
		{"edit product", "/api/v1/product/edit", productEditOutcomeJSON},
		{"remove product", "/api/v1/product/delete", productToDeleteJSON},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("try %s with incorrect http method type", tc.description), func(t *testing.T) {
			productsBeforeApiCall := dbStub.Products
			response := server.callApi(http.MethodPatch, tc.apiPath, tc.json)
			checkStatus(t, http.StatusMethodNotAllowed, response.Code)
			checkProducts(t, productsBeforeApiCall, server.store.GetAllProducts())
		})
	}
}
