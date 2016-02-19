package handlers

import (
  "github.com/tksasha/rest"

  . "github.com/tksasha/balance/models"
)

var Recoveries *rest.Handler

func init() {
  Recoveries = &rest.Handler{
    Resource: func() rest.Resource { return new(Recovery) },
  }
}
