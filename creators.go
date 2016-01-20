package main

import (
  "net/url"
)

func CreateCategory(params url.Values) (Category, Errors) {
  errors := make(Errors)

  var category Category

  name := params.Get("category[name]")

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
