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
  router.GET("/categories", Categories.Index)

  router.POST("/categories", Categories.Create)

  //router.POST("/categories/:id/recovery", new(Categories).Recovery)

  router.PATCH("/categories/:id", Categories.Update)

  router.PUT("/categories/:id", Categories.Update)

  router.DELETE("/categories/:id", Categories.Destroy)

  //
  // /items
  //
  router.GET("/items", Items.Index)

  router.POST("/items", Items.Create)

  router.DELETE("/items/:id", Items.Destroy)

  router.PATCH("/items/:id", Items.Update)

  router.PUT("/items/:id", Items.Update)

  //
  // /consolidates
  //
  //router.GET("/consolidates", new(Consolidates).Index)

  //
  // /cashes
  //
  router.POST("/cashes", Cashes.Create)

  router.GET("/cashes", Cashes.Index)

  return alice.New(LoggingHandler, ContentTypeHandler).Then(router)
}
