package handlers

import (
  . "github.com/tksasha/balance/rest/handler"
  . "github.com/tksasha/balance/rest/resource"
  . "github.com/tksasha/balance/rest/collection"
  . "github.com/tksasha/balance/rest/router"
  . "github.com/tksasha/balance/models"
)

var Categories *Handler

func init() {
  Categories = &Handler{
    Collection: func() Collection { return new(CategoryCollection) },
    Resource:   func() Resource { return new(Category) },
  }

  Router.GET("/categories", Categories.Index)

  Router.POST("/categories", Categories.Create)

  Router.PATCH("/categories/:id", Categories.Update)

  Router.PUT("/categories/:id", Categories.Update)

  Router.DELETE("/categories/:id", Categories.Destroy)
}
