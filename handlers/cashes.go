package handlers

import (
  . "../models"
)

var Cashes *Handler

func init() {
  Cashes = &Handler {
    Collection: func()  Collection  { return new(CashCollection) },
    Resource:   func()  Resource    { return new(Cash) },
  }
}
