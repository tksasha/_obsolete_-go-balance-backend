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

  router.GET("/items", ItemsIndex)

  router.POST("/items", ItemsCreate)

  router.GET("/consolidates", ConsolidatesIndex)

  return alice.New(LoggingHandler, ContentTypeHandler).Then(router)
}
