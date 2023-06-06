package server

import (
	"encoding/json"
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
func (s Server) callApi(method, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	response := httptest.NewRecorder()

	s.ServeHTTP(response, request)
	return response
}

func TestServer(t *testing.T) {
	server := New()

	t.Run("GET all products", func(t *testing.T) {
		response := server.callApi(http.MethodGet, "/api/v1/products/all")

		checkStatus(t, http.StatusOK, response.Code)

		var got []Product

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of products, '%v'", response.Body, err)
		}

		want := []Product{
			{"juice"},
			{"cheese"},
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("POST new product", func(t *testing.T) {
		response := server.callApi(http.MethodPost, "/api/v1/product/new")

		checkStatus(t, http.StatusAccepted, response.Code)
	})

}
