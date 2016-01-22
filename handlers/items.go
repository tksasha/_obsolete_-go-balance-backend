package handlers

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/tksasha/go-date"

  . "../config"
  . "../models"
)

type Items int

func (Items) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  params := r.URL.Query()

  d := date.New(params.Get("year"), params.Get("month"))

  var items []Item

  DB.
    Where("date BETWEEN ? AND ?", d.BeginningOfMonth().String(), d.EndOfMonth().String()).
    Order("date").
    Preload("Category").
    Find(&items)

  render(w, &items)
}

func (Items) Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  id := params.ByName("id")

  var item Item

  if DB.First(&item, id).RecordNotFound() {
    http.NotFound(w, r)
  } else {
    render(w, &item)
  }
}

func (Items) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  item := Item{}

  item.Build(r.Form)

  render(w, &item)
}
