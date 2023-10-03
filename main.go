package main

import (
	"log"
	"net/http"

	"github.com/paulina-sorys/shop-organiser-api/db"
	"github.com/paulina-sorys/shop-organiser-api/model"
	"github.com/paulina-sorys/shop-organiser-api/server"
)

func main() {
	db := &db.InMemeoryDB{Products: []model.Product{{Name: "test"}}}
	log.Fatal(http.ListenAndServe(":5000", server.New(db).Handler))
}
