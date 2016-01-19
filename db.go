package main

import (
  "log"

  _ "github.com/mattn/go-sqlite3"
  "github.com/jinzhu/gorm"
)

func NewDBConnection() (db gorm.DB) {
  var err error

  if db, err = gorm.Open("sqlite3", "/Users/tksasha/balance/db/development.sqlite3"); err != nil {
    log.Fatalln(err)
  }

  return
}

//
// Scopes
//
func Visible(db *gorm.DB) *gorm.DB {
  return db.Where("visible = 't'")
}
