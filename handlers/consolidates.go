package handlers

import (
  . "./rest"
  . "../models"
)

var Consolidates *RESTHandler

func init() {
  Consolidates = &RESTHandler {
    Collection: func() Collection { return new(ConsolidateCollection) },
  }
}
