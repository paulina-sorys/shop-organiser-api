package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETProductsAll(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "api/v1/products/all", nil)
	response := httptest.NewRecorder()

	New(response, request)

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
}
