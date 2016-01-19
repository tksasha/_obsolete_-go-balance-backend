package main

import (
	"log"
	"net/http"

  "github.com/jinzhu/gorm"
)

var db gorm.DB

func main() {
  db = NewDBConnection()

  defer db.Close()

  db.LogMode(true)

	log.Fatalln(http.ListenAndServe(":3000", NewRouter()))
}
