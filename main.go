package main

import (
  "log"
  "net/http"

  . "./config"
  _ "./config/development"
  . "./router"
)

func main() {
  defer DB.Close()

  //
  // Serve HTTP Queries
  //
  log.Fatalln(http.ListenAndServe(":3000", NewRouter()))
}
