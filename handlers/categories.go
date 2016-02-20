package handlers

import (
  . "github.com/tksasha/balance/rest/handler"
  . "github.com/tksasha/balance/models"
)

var Categories *Handler

func init() {
  Categories = &Handler{
    Collection: func() Collection { return new(CategoryCollection) },
    Resource:   func() Resource { return new(Category) },
  }
}
