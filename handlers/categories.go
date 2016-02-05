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

//
// POST /categories
//
func (h Categories) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  category := new(Category)

  if category.IsCreate(r.Form) {
    h.render(w, category, 200)
  } else {
    h.render(w, category.Errors, 422)
  }
}

//
// PATCH /categories
//
func (h Categories) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  r.ParseForm()

  category := new(Category)

  if DB.First(category, params.ByName("id")).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  if category.IsUpdate(r.Form) {
    h.render(w, category, 200)
  } else {
    h.render(w, category.Errors, 422)
  }
}

//
// DELETE /categories
//
func(h Categories) Destroy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  category := new(Category)

  if DB.First(category, params.ByName("id")).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  DB.Delete(category)

  h.render(w, "OK", 200)
}

//
// POST /categories/:id/recovery
//
func(h Categories) Recovery(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  category := new(Category)

  if DB.Unscoped().First(category, params.ByName("id")).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  DB.Unscoped().Model(&category).Update("deleted_at", nil)

  h.render(w, category, 200)
}
