package middleware

import (
  "net/http"
  "os"

  "github.com/gorilla/handlers"
)

func LogMiddleware(h http.Handler) http.Handler {
  return handlers.LoggingHandler(os.Stdout, h)
}
