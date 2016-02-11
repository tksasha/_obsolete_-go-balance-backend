package router

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/justinas/alice"

  . "../handlers"
  . "../handlers/middleware"
)

func NewRouter() http.Handler {
  router := httprouter.New()

  //
  // /categories
  //
  router.GET("/categories", Categories.Index)

  router.POST("/categories", Categories.Create)

  router.POST("/categories/:category_id/recovery", Recoveries.Create)

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
  router.GET("/consolidates", Consolidates.Index)

  //
  // /cashes
  //
  router.POST("/cashes", Cashes.Create)

  router.GET("/cashes", Cashes.Index)

  router.DELETE("/cashes/:id", Cashes.Destroy)

  return alice.New(LogMiddleware, ContentTypeMiddleware).Then(router)
}
