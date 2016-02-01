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

func (h Categories) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var categories []Category

  DB.
    Scopes(VisibleCategories).
    Order("income").
    Find(&categories)

  h.render(w, &categories)
}

func (h Categories) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  category := new(Category)

  if category.IsCreate(r.Form) {
    h.render(w, category)
  } else {
    h.render(w, category.Errors)
  }
}

func (h Categories) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  r.ParseForm()

  category := new(Category)

  if DB.Where("id = ?", params.ByName("id")).First(category).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  if category.IsUpdate(r.Form) {
    h.render(w, category)
  } else {
    h.render(w, category.Errors)
  }
}
