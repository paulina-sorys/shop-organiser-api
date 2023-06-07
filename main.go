package main

import (
	"log"
	"net/http"

	"github.com/paulina-sorys/shop-organiser/db"
	"github.com/paulina-sorys/shop-organiser/model"
	"github.com/paulina-sorys/shop-organiser/server"
)

func main() {
	db := &db.InMemeoryDB{Products: []model.Product{{Name: "test"}}}
	log.Fatal(http.ListenAndServe(":5000", server.New(db).Handler))
}
