package handlers

import (
  . "./rest"
  . "../models"
)

var Recoveries *RESTHandler

func init() {
  Recoveries = &RESTHandler {
    Resource: func() Resource { return new(Recovery) },
  }
}
