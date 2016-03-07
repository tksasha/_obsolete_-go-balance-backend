package models

import (
  "net/url"
  "strings"
  "time"

  . "github.com/tksasha/balance/rest/model"
  . "github.com/tksasha/balance/rest/db"
  _ "github.com/tksasha/balance/config"
)

func init() {
  DB.AutoMigrate(&Category{})
}

type Category struct {
  Model

  ID        int        `json:"id"`
  Name      string     `json:"name"     valid:"present"`
  Income    bool       `json:"income"`
  CreatedAt time.Time  `json:"-"`
  UpdatedAt time.Time  `json:"-"`
  DeletedAt *time.Time `json:"-"`
}

type CategoryCollection []Category

//
// Category.Build
//
func (c *Category) Build(values url.Values) {
  if name := values.Get("category[name]"); name != "" {
    c.Name = name
  }

  if income := values.Get("category[income]"); income != "" {
    switch values.Get("category[income]") {
    case "1", "true":
      c.Income = true
    default:
      c.Income = false
    }
  }
}

//
// Category.Update
//
func (c *Category) Update() {
  DB.Save(c)
}

//
// DEPRECATED
//
func (c *Category) IsValid() bool {
  c.validateUniquenessOfName()

  return c.Errors().IsEmpty()
}

//
// DEPRECATED
//
func (c *Category) validateUniquenessOfName() {
  var count int

  query := DB.Table("categories").Where("LOWER(name)=?", strings.ToLower(c.Name))

  if c.ID != 0 {
    query = query.Not("id = ?", c.ID)
  }

  query.Count(&count)

  if count > 0 {
    c.Errors().Add("name", "already exists")
  }
}

//
// CategoryCollection.Search
//
func (c *CategoryCollection) Search(values url.Values) {
  DB.Order("income").Find(&c)
}
