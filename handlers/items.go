package handlers

import (
  "net/http"

  "github.com/julienschmidt/httprouter"

  . "../config"
  . "../models"
)

type Items struct {
  BaseHandler
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

//
// DELETE /items
//
func (i Items) Destroy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  item := new(Item)

  if DB.First(item, params.ByName("id")).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  DB.Delete(item)

  i.render(w, "OK", 200)
}

//
// PATCH /items
//
func (i Items) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  r.ParseForm()

  item := new(Item)

  if DB.First(item, params.ByName("id")).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  if item.IsUpdate(r.Form) {
    i.render(w, item, 200)
  } else {
    i.render(w, item.Errors, 422)
  }
}
