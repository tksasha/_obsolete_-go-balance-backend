package handlers

import (
  "github.com/tksasha/rest"

  . "github.com/tksasha/balance/models"
)

var Cashes *rest.Handler

func init() {
  Cashes = &rest.Handler {
    Collection: func()  rest.Collection  { return new(CashCollection) },
    Resource:   func()  rest.Resource    { return new(Cash) },
  }
}
