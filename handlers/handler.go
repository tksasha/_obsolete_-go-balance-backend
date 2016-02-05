package handlers

import (
  "net/http"
  "encoding/json"

  "github.com/julienschmidt/httprouter"

  . "../config"
  . "../models"
)

type Handler struct {
  Resource    func() Resource
  Collection  func() Collection
}

//
// GET - Show Resource
//

//
// GET - List Collection
//
func (h Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  collection := h.Collection()

  collection.Search(r.URL.Query())

  h.render(w, collection, 200)
}

//
// POST - Create Resource
//
func (h Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  resource := h.Resource()

  if resource.IsCreate(r.Form) {
    h.render(w, resource, 200)
  } else {
    h.render(w, resource.GetErrors(), 422)
  }
}

//
// DELETE - Destroy Resource
//
func (h Handler) Destroy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  resource := h.Resource()

  if DB.First(resource, params.ByName("id")).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  DB.Delete(resource)

  h.render(w, "OK", 200)
}

//
// PATCH - Update Resource
//
func(h Handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  r.ParseForm()

  resource := h.Resource()

  if DB.First(resource, params.ByName("id")).RecordNotFound() {
    http.NotFound(w, r)

    return
  }

  if resource.IsUpdate(r.Form) {
    h.render(w, resource, 200)
  } else {
    h.render(w, resource.GetErrors(), 422)
  }
}

//
// Service Methods
//
func (Handler) render(w http.ResponseWriter, items interface{}, code int) {
  w.WriteHeader(code)

  if err := json.NewEncoder(w).Encode(items); err != nil {
    panic(err)
  }
}
