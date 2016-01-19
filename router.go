package main

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/justinas/alice"
)

func NewRouter() http.Handler {
  router := httprouter.New()

  for _, route := range routes {
    router.Handle(route.Method, route.Path, route.HandlerFunc)
  }

  return alice.New(LoggingHandler, ContentTypeHandler).Then(router)
}
