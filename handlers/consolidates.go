package handlers

import (
  . "github.com/tksasha/balance/rest/handler"
  . "github.com/tksasha/balance/rest/collection"
  . "github.com/tksasha/balance/models"
)

var Consolidates *Handler

func init() {
  Consolidates = &Handler{
    Collection: func() Collection { return new(ConsolidateCollection) },
  }
}
