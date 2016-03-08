package test

import (
  "net/url"
  "log"
  "os"

  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"

  . "github.com/tksasha/balance/rest/db"
)

var Values url.Values

func init() {
  Values = make(url.Values)

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
}
