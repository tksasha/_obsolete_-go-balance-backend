package handlers

import (
  . "github.com/tksasha/balance/rest/handler"
  . "github.com/tksasha/balance/rest/resource"
  . "github.com/tksasha/balance/rest/collection"
  . "github.com/tksasha/balance/rest/router"
  . "github.com/tksasha/balance/models"
)

var Cashes *Handler

func init() {
  Cashes = &Handler{
    Collection: func() Collection { return new(CashCollection) },
    Resource:   func() Resource { return new(Cash) },
  }

  Router.POST("/cashes", Cashes.Create)

  Router.GET("/cashes", Cashes.Index)

  Router.DELETE("/cashes/:id", Cashes.Destroy)

  Router.PATCH("/cashes/:id", Cashes.Update)
}
