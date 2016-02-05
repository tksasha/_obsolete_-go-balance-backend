package handlers

import (
  . "../models"
)

var Categories *Handler

func init() {
  Categories = &Handler {
    Collection: func() Collection { return new(CategoryCollection) },
    Resource:   func() Resource   { return new(Category) },
  }
}
