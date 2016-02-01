package handlers

import (
  "net/http"

  "github.com/julienschmidt/httprouter"

  . "../config"
  . "../models"
)

type Categories struct {
  BaseHandler
}

func (c Categories) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var categories []Category

  DB.
    Scopes(VisibleCategories).
    Order("income").
    Find(&categories)

  c.render(w, &categories)
}

func (c Categories) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  category := Category{}

  if isValid, errors := category.Create(r.Form); isValid == true {
    c.render(w, &category)
  } else {
    c.render(w, &errors)
  }
}
