package main

import (
  "net/url"
)


type Category struct {
  ID      int     `json:"id"`
  Name    string  `json:"name"`
  Income  bool    `json:"income"`
  Visible bool    `json:"visible"`
}

func CreateCategory(params url.Values) (Category, map[string][]string) {
  errs := make(map[string][]string)

  var category Category

  name := ""

  if params["category[name]"] != nil {
    name = params.Get("category[name]")
  }

  if name == "" {
    errs["name"] = append(errs["Name"], "can't be blank")
  }

  if db.Where("name=?", name).First(&category).RecordNotFound() {
    // do nothing
  } else {
    errs["Name"] = append(errs["Name"], "already exists")
  }

  if len(errs["name"]) == 0 {
    category.Name = name

    category.Visible = true

    db.Create(&category)
  }

  return category, errs
}
