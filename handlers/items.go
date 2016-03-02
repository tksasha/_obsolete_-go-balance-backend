package handlers

import (
  . "github.com/tksasha/balance/rest/handler"
  . "github.com/tksasha/balance/rest/resource"
  . "github.com/tksasha/balance/rest/collection"
  . "github.com/tksasha/balance/rest/router"
  . "github.com/tksasha/balance/models"
)

var Items *Handler

func init() {
  Items = &Handler{
    Collection: func() Collection { return new(ItemCollection) },
    Resource:   func() Resource { return new(Item) },
  }

  Router.GET("/items", Items.Index)

  Router.POST("/items", Items.Create)

  Router.DELETE("/items/:id", Items.Delete)

  Router.PATCH("/items/:id", Items.Update)

  Router.PUT("/items/:id", Items.Update)
}
