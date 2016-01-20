package main

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/justinas/alice"
)

func NewRouter() http.Handler {
  router := httprouter.New()

  router.GET("/categories", CategoriesIndex)

  router.POST("/categories", CategoriesCreate)

  return alice.New(LoggingHandler, ContentTypeHandler).Then(router)
}
