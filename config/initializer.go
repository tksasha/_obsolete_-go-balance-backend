package config

import (
  "log"
  "os"
  "time"

  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

var DB gorm.DB

func init() {
  //
  // Set TimeZone
  //
  time.Local = time.UTC

  //
  // Check $DBURL
  //
  dburl := os.Getenv("DBURL")

  if dburl == "" {
    log.Fatalln("$DBURL must be specified")
  }

  var err error

  //
  // Make DataBase Connection
  //
  if DB, err = gorm.Open("postgres", dburl); err != nil {
    panic(err)
  }

  //
  // Show SQL in LOG
  //
  DB.LogMode(false)
}
