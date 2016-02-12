package handlers

import (
  "github.com/tksasha/go-rest"

  . "../models"
)

var Cashes *rest.Handler

func init() {
  Cashes = &rest.Handler {
    Collection: func()  rest.Collection  { return new(CashCollection) },
    Resource:   func()  rest.Resource    { return new(Cash) },
  }
}
