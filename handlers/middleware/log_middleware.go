package middleware

import (
	"net/http"
	"os"

	"github.com/tksasha/balance/Godeps/_workspace/src/github.com/gorilla/handlers"
)

func LogMiddleware(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
