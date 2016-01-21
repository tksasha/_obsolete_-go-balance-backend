package main

import (
  . "github.com/tksasha/go-errors"
)

//
// Category
//
func (c *Category) Valid() (bool, Errors) {
  errors := make(Errors)

  if c.Name == "" {
    errors.Add("name", "can't be blank")
  }

  var count int

  db.Table("categories").Where("name=?", c.Name).Count(&count)

  if count > 0 {
    errors.Add("name", "already exists")
  }

  return errors.IsEmpty(), errors
}
