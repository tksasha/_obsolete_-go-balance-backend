package handlers

import (
  "github.com/tksasha/rest"

  . "github.com/tksasha/balance/models"
)

var Items *rest.Handler

func init() {
  Items = &rest.Handler {
    Collection: func() rest.Collection { return new(ItemCollection) },
    Resource:   func() rest.Resource   { return new(Item) },
  }
}
