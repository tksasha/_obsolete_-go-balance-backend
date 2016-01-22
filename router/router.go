package router

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/justinas/alice"

  . "../handlers"
)

func NewRouter() http.Handler {
  router := httprouter.New()

  router.GET("/categories", new(Categories).Index)

  router.POST("/categories", new(Categories).Create)

  router.GET("/items", new(Items).Index)

  router.POST("/items", new(Items).Create)

  router.GET("/consolidates", new(Consolidates).Index)

  return alice.New(LoggingHandler, ContentTypeHandler).Then(router)
}
