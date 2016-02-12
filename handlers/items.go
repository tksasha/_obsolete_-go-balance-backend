package handlers

import (
  "github.com/tksasha/go-rest"

  . "../models"
)

var Items *rest.Handler

func init() {
  Items = &rest.Handler {
    Collection: func() rest.Collection { return new(ItemCollection) },
    Resource:   func() rest.Resource   { return new(Item) },
  }
}
