package models

import (
  "net/url"

  . "../config"
)

type Category struct {
  BaseModel

  ID      int     `json:"id"`
  Name    string  `json:"name"`
  Income  bool    `json:"income"`
  Visible bool    `json:"visible" sql:"default:'t'"`
}

func (c *Category) Build(values url.Values) {
  c.Init()

  c.Name = values.Get("category[name]")
}

func (c *Category) IsValid() bool {
  c.validatePresenceOfName()

  c.validateUniquenessOfName()

  return c.Errors.IsEmpty()
}

func (c *Category) IsCreate(values url.Values) bool {
  c.Build(values)

  if c.IsValid() {
    DB.Create(c)

    return true
  } else {
    return false
  }
}

func (c *Category) IsUpdate(values url.Values) bool {
  c.Build(values)

  if c.IsValid() {
    DB.Save(c)

    return true
  } else {
    return false
  }
}

func (c *Category) validatePresenceOfName() {
  if c.Name == "" {
    c.Errors.Add("name", "can't be blank")
  }
}

func (c *Category) validateUniquenessOfName() {
  var count int

  query := DB.Table("categories").Where("name=?", c.Name)

  if c.ID != 0 {
    query = query.Not("id = ?", c.ID)
  }

  query.Count(&count)

  if count > 0 {
    c.Errors.Add("name", "already exists")
  }
}
