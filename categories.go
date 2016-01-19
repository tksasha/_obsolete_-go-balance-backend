package main

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func CategoriesIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var categories []Category

  db.
    Scopes(VisibleCategories).
    Order("income").
    Find(&categories)

  render(w, &categories)
}

func CategoriesCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  if category, err := CreateCategory(r.Form); err != nil {
    render(w, &err)
  } else {
    render(w, &category)
  }
}
