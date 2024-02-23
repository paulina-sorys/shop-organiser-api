package main

import (
	"log"
	"net/http"

	"db"
	"model"
	"server"
)

func main() {
	db := &db.InMemoryDB{Products: []model.Product{{Name: "test"}}}
	log.Fatal(http.ListenAndServe(":5000", server.New(db).Handler))
}
