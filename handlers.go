package main

import (
  "net/http"
  "encoding/json"
  "log"

  "github.com/julienschmidt/httprouter"
)

//
// GET /items
//
func ItemsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  params := r.URL.Query()

  t := NewTime(params["year"][0], params["month"][0])

  var items []Item

  db.
    Where("date BETWEEN ? AND ?", t.BeginningOfMonth().UTC(), t.EndOfMonth().UTC()).
    Order("date").
    Find(&items)

  render(w, &items)
}

//
// GET /items/:id
//
func ItemsShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  id := params.ByName("id")

  var item Item

  if db.First(&item, id).RecordNotFound() {
    http.NotFound(w, r)
  } else {
    render(w, &item)
  }
}

//
// POST /items
//
func ItemsCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  // Do nothing
}

//
// GET /categories
//
func CategoriesIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var categories []Category

  db.
    Scopes(VisibleCategories).
    Order("income").
    Find(&categories)

  render(w, &categories)
}

//
// POST /categories
//
func CategoriesCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  if category, err := CreateCategory(r.Form); err != nil {
    render(w, &err)
  } else {
    render(w, &category)
  }
}

//
// GET /consolidates
//
func ConsolidatesIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var consolidates []Consolidate

  db.
    Table("items").
    Select("SUM(sum) AS sum, category_id").
    Group("category_id").
    Scan(&consolidates)

  render(w, consolidates)
}

//
// Service Handlers
//
func ContentTypeHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    h.ServeHTTP(w, r)
  })
}

func LoggingHandler(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s", r.Method, r.URL)

    h.ServeHTTP(w, r)
  })
}

func render(w http.ResponseWriter, items interface{}) {
  if err := json.NewEncoder(w).Encode(items); err != nil {
    log.Fatalln(err)
  }
}
