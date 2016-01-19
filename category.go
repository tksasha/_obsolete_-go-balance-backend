package main

import (
  "net/url"

  "github.com/jinzhu/gorm"
)

type Category struct {
  ID      int     `json:"id"`
  Name    string  `json:"name"`
  Income  bool    `json:"income"`
  Visible bool    `json:"visible"`
}

func CreateCategory(params url.Values) (Category, Errors) {
  errors := make(Errors)

  var category Category

  name := ""

  if params["category[name]"] != nil {
    name = params.Get("category[name]")
  }

  if name == "" {
    errors.Set("name", "can't be blank")
  }

  if !db.Where("name=?", name).First(&category).RecordNotFound() {
    errors.Set("name", "already exists")
  }

  if errors.IsEmpty() {
    category.Name = name

    category.Visible = true

    db.Create(&category)

    return category, nil
  } else {
    return category, errors
  }
}

//
// Scopes
//
func VisibleCategories(db *gorm.DB) *gorm.DB {
  return db.Where("visible IN('t', 1)")
}
