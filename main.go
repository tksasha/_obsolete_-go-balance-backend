package main

import (
  "log"
  "net/http"
  "os"

  . "github.com/tksasha/balance/rest/router"
  . "github.com/tksasha/balance/rest/db"
  _ "github.com/tksasha/balance/config"
  _ "github.com/tksasha/balance/handlers"
)

func main() {
  defer DB.Close()

  port := os.Getenv("PORT")

  if port == "" {
    log.Fatalln("$PORT must be specified")
  }

  //
  // Serve HTTP Queries
  //
  log.Fatalln(http.ListenAndServe(":" + port, Router))
}
