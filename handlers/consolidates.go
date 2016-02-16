package handlers

import (
  "github.com/tksasha/rest"

  . "github.com/tksasha/balance/models"
)

var Consolidates *rest.Handler

func init() {
  Consolidates = &rest.Handler {
    Collection: func() rest.Collection { return new(ConsolidateCollection) },
  }
}
