package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETProductsAll(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "api/v1/products/all", nil)
	response := httptest.NewRecorder()

	New(response, request)

	got := response.Body.String()
	want := "20"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
