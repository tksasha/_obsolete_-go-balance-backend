package handlers

import (
  . "github.com/tksasha/balance/rest/handler"
  . "github.com/tksasha/balance/rest/resource"
  . "github.com/tksasha/balance/models"
)

var Recoveries *Handler

func init() {
  Recoveries = &Handler{
    Resource: func() Resource { return new(Recovery) },
  }
}
