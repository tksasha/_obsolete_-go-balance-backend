package handlers

import (
  "github.com/tksasha/rest"

  . "github.com/tksasha/balance/models"
)

var Categories *rest.Handler

func init() {
  Categories = &rest.Handler{
    Collection: func() rest.Collection { return new(CategoryCollection) },
    Resource:   func() rest.Resource { return new(Category) },
  }
}
