package main

import (
  "net/http"

  "github.com/jinzhu/gorm"
  "github.com/julienschmidt/httprouter"
)

func CategoriesIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var categories []Category

  db.Scopes(VisibleCategories).Order("income").Find(&categories)

  render(w, categories)
}

func CategoriesCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  r.ParseForm()

  _, err := CreateCategory(r.Form)

  render(w, err)
}

//
// Scopes
//
func VisibleCategories(db *gorm.DB) *gorm.DB {
  return db.Where("visible = 't'")
}
