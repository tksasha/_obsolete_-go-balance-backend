package handlers

import (
  . "../models"
)

var Items *Handler

func init() {
  Items = &Handler {
    Collection: func() Collection { return new(ItemCollection) },
    Resource:   func() Resource   { return new(Item) },
  }
}
