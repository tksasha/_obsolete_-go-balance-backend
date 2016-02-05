package handlers

import (
  "net/http"
  "encoding/json"

  "github.com/julienschmidt/httprouter"

  . "../models"
)

type Handler struct {
  Resource    Resource
  Collection  Collection
}

func (h Handler) Show(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  First(h.Resource)

  h.render(w, h.Resource, 200)
}

func (h Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  h.Collection.Search(r.URL.Query())

  h.render(w, h.Collection, 200)
}

func (Handler) render(w http.ResponseWriter, items interface{}, code int) {
  w.WriteHeader(code)

  if err := json.NewEncoder(w).Encode(items); err != nil {
    panic(err)
  }
}
