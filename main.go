package main

import (
  "log"
  "net/http"

  . "github.com/tksasha/balance/config"
  _ "github.com/tksasha/balance/config/development"
  . "github.com/tksasha/balance/router"
)

func main() {
  defer DB.Close()

  //
  // Serve HTTP Queries
  //
  log.Fatalln(http.ListenAndServe(":3000", NewRouter()))
}
