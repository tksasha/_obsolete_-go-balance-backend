package handlers

import (
  . "github.com/tksasha/balance/rest/handler"
  . "github.com/tksasha/balance/models"
)

var Cashes *Handler

func init() {
  Cashes = &Handler{
    Collection: func() Collection { return new(CashCollection) },
    Resource:   func() Resource { return new(Cash) },
  }
}
