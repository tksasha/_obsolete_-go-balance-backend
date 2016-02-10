package handlers

import (
  . "../models"
)

var Consolidates *Handler

func init() {
  Consolidates = &Handler {
    Collection: func() Collection { return new(ConsolidateCollection) },
  }
}
