package main

import (
  "log"
  "net/http"
  "fmt"
  "os"
  "time"

  _ "github.com/mattn/go-sqlite3"
  "github.com/jinzhu/gorm"

  . "./config"
  . "./models"
  . "./router"
)

func main() {
  var err error

  //
  // Set TimeZone
  //
  time.Local = time.UTC

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

  if DB, err = gorm.Open("sqlite3", url); err != nil {
    log.Fatalln(err)
  }

  defer DB.Close()

  //
  // Run Migrations
  //
  DB.AutoMigrate(&Item{}, &Category{})

  //
  // Show SQL in LOG
  //
  DB.LogMode(true)

  //
  // Serve HTTP Queries
  //
  log.Fatalln(http.ListenAndServe(":3000", NewRouter()))
}
