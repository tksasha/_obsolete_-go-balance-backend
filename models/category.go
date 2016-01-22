package models

import (
  "net/url"

  . "github.com/tksasha/go-errors"

  . "../config"
)

type Category struct {
  ID      int     `json:"id"`
  Name    string  `json:"name"`
  Income  bool    `json:"income"`
  Visible bool    `json:"visible" sql:"default:'t'"`
}

func (c *Category) Build(values url.Values) {
  c.Name = values.Get("category[name]")
}

func (c *Category) Valid() (bool, Errors) {
  errors := make(Errors)

  if c.Name == "" {
    errors.Add("name", "can't be blank")
  }

  var count int

  DB.Table("categories").Where("name=?", c.Name).Count(&count)

  if count > 0 {
    errors.Add("name", "already exists")
  }

  return errors.IsEmpty(), errors
}

func (c *Category) Create(values url.Values) (bool, Errors) {
  c.Build(values)

  if isValid, errors := c.Valid(); isValid == true {
    DB.Create(&c)

    return true, nil
  } else {
    return false, errors
  }
}
