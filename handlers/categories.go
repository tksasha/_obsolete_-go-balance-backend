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
// GET /categories
//
func (h Categories) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var categories []Category

  DB.
    Order("income").
    Find(&categories)

  h.render(w, &categories)
}

//
// POST /categories
//
func (h Categories) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  category := new(Category)

  if category.IsCreate(r.Form) {
    h.render(w, category)
  } else {
    h.render(w, category.Errors, StatusUnprocessableEntity)
  }
}

//
// PATCH /categories
//
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

//
// DELETE /categories
//
func(h Categories) Destroy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  category := new(Category)

  if DB.Where("id = ?", params.ByName("id")).First(category).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  DB.Delete(category)

  h.render(w, "OK")
}

//
// POST /categories/:id/recovery
//
func(h Categories) Recovery(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  category := new(Category)

  if DB.Unscoped().Where("id = ?", params.ByName("id")).First(category).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  DB.Unscoped().Model(&category).Update("deleted_at", nil)

  h.render(w, category)
}
