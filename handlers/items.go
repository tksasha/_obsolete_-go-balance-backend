package handlers

import (
  . "../models"
)

var Items *RESTHandler

func init() {
  Items = &RESTHandler {
    Collection: func() Collection { return new(ItemCollection) },
    Resource:   func() Resource   { return new(Item) },
  }
}
