package handlers

import (
  "net/http"
  "log"
)

func LoggingHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s", r.Method, r.URL)

    h.ServeHTTP(w, r)
  })
}
