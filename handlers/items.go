package handlers

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/tksasha/go-date"

  . "../config"
  . "../models"
)

type Items struct {
  BaseHandler
}

//
// GET /items
//
func (i Items) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  params := r.URL.Query()

  d := date.New(params.Get("year"), params.Get("month"))

  var items []Item

  DB.
    Where("date BETWEEN ? AND ?", d.BeginningOfMonth().String(), d.EndOfMonth().String()).
    Order("date").
    Preload("Category").
    Find(&items)

  i.render(w, items, 200)
}

//
// GET /items/:id
//
func (i Items) Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  id := params.ByName("id")

  var item Item

  if DB.First(&item, id).RecordNotFound() {
    http.NotFound(w, r)
  } else {
    i.render(w, item, 200)
  }
}

//
// POST /items
//
func (h Items) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  item := new(Item)

  if item.IsCreate(r.Form) {
    h.render(w, item, 200)
  } else {
    h.render(w, item.Errors, 422)
  }
}
