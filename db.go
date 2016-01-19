package main

import (
  "log"
  "fmt"
  "os"

  _ "github.com/mattn/go-sqlite3"
  "github.com/jinzhu/gorm"
)

func NewDBConnection() (db gorm.DB) {
  var err error

  //
  // TODO: move it to main()
  //
  dir := fmt.Sprintf("%s/.balance/db",  os.Getenv("HOME"))

  if err = os.MkdirAll(dir, 0755); err != nil {
    panic(err)
  }

  url := fmt.Sprintf("%s/development.sqlite3", dir)

  if db, err = gorm.Open("sqlite3", url); err != nil {
    log.Fatalln(err)
  }

  db.AutoMigrate(&Item{}, &Category{})

  return
}
