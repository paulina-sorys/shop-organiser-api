package main

import (
	"log"
	"net/http"

	"github.com/paulina-sorys/shop-organiser/server"
)

func main() {
	handler := http.HandlerFunc(server.New)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
