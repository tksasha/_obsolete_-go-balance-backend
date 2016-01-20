package main

import (
  "net/url"
)

//
// Category
//
func (c *Category) Create(values url.Values) (bool, Errors) {
  c.Build(values)

  if isValid, errors := c.Valid(); isValid == true {
    db.Create(&c)

    return true, nil
  } else {
    return false, errors
  }
}
