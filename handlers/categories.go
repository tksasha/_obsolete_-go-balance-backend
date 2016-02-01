package handlers

import (
  "net/http"

  "github.com/julienschmidt/httprouter"

  . "../config"
  . "../models"
)

type Categories BaseHandler

func (Categories) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var categories []Category

  DB.
    Scopes(VisibleCategories).
    Order("income").
    Find(&categories)

  render(w, &categories)
}

func (Categories) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  category := Category{}

  if isValid, errors := category.Create(r.Form); isValid == true {
    render(w, &category)
  } else {
    render(w, &errors)
  }
}
