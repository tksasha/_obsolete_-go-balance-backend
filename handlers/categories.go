package handlers

import (
  . "../models"
)

var Categories *RESTHandler

func init() {
  Categories = &RESTHandler {
    Collection: func() Collection { return new(CategoryCollection) },
    Resource:   func() Resource   { return new(Category) },
  }
}
