package handlers

import (
  . "../models"
)

var Consolidates *RESTHandler

func init() {
  Consolidates = &RESTHandler {
    Collection: func() Collection { return new(ConsolidateCollection) },
  }
}
