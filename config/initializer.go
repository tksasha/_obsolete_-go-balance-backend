package config

import (
  "time"
  "os"
  "log"

  _ "github.com/lib/pq"
  "github.com/jinzhu/gorm"
)

var DB gorm.DB

func init() {
  //
  // Set TimeZone
  //
  time.Local = time.UTC

  DBURL := os.Getenv("DBURL")

  if DBURL == "" {
    log.Fatalln("$DBURL must be specified")
  }

  var err error

  //
  // Make DataBase Connection
  //
  if DB, err = gorm.Open("postgres", DBURL); err != nil {
    panic(err)
  }

  //
  // Show SQL in LOG
  //
  DB.LogMode(true)
}
