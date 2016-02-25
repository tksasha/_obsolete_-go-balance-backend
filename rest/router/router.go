package router

import (
  "github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

func init() {
  Router = httprouter.New()
}
