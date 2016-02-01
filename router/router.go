package router

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/justinas/alice"

  . "../handlers"
)

func NewRouter() http.Handler {
  router := httprouter.New()

  //
  // /categories
  //
  router.GET("/categories", new(Categories).Index)

  router.POST("/categories", new(Categories).Create)

  router.PATCH("/categories/:id", new(Categories).Update)

  //
  // /items
  //
  router.GET("/items", new(Items).Index)

  router.POST("/items", new(Items).Create)

  //
  // /consolidates
  //
  router.GET("/consolidates", new(Consolidates).Index)

  return alice.New(LoggingHandler, ContentTypeHandler).Then(router)
}
