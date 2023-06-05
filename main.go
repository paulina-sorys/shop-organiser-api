package main

import (
	"log"
	"net/http"

	"github.com/paulina-sorys/shop-organiser/server"
)

func main() {
	server := server.New()
	log.Fatal(http.ListenAndServe(":5000", server.Handler))
}
