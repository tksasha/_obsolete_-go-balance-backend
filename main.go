package main

import (
  "log"
  "net/http"
  "fmt"
  "os"

  _ "github.com/mattn/go-sqlite3"
  "github.com/jinzhu/gorm"
)

var db gorm.DB

func main() {
  var err error

  //
  // Prepare Workdir
  //
  dir := fmt.Sprintf("%s/.balance/db",  os.Getenv("HOME"))

  if err = os.MkdirAll(dir, 0755); err != nil {
    panic(err)
  }

  //
  // Make DataBase Connection
  //
  url := fmt.Sprintf("%s/development.sqlite3", dir)

  if db, err = gorm.Open("sqlite3", url); err != nil {
    log.Fatalln(err)
  }

  defer db.Close()

  //
  // Run Migrations
  //
  db.AutoMigrate(&Item{}, &Category{})

  //
  // Show SQL in LOG
  //
  db.LogMode(true)

  //
  // Serve HTTP Queries
  //
	log.Fatalln(http.ListenAndServe(":3000", NewRouter()))
}
