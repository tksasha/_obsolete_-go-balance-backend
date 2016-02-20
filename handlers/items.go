package handlers

import (
  . "github.com/tksasha/balance/rest/handler"
  . "github.com/tksasha/balance/models"
)

var Items *Handler

func init() {
  Items = &Handler{
    Collection: func() Collection { return new(ItemCollection) },
    Resource:   func() Resource { return new(Item) },
  }
}
