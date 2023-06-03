package server

import (
	"fmt"
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
