package models

import (
  "net/url"
  "strconv"
  "strings"

  . "github.com/tksasha/balance/rest/model"
  . "github.com/tksasha/balance/rest/db"
  _ "github.com/tksasha/balance/config"
)

func init() {
  DB.AutoMigrate(&Cash{})
}

type Cash struct {
  Model

  ID   int
  Name string   `valid:"present"`
  Sum  float64
}

type CashCollection []Cash

//
// Cash.Build
//
func (c *Cash) Build(values url.Values) {
  if len(values["cash[name]"]) != 0 {
    c.Name = values.Get("cash[name]")
  }

  if len(values["cash[sum]"]) != 0 {
    if sum, err := strconv.ParseFloat(values.Get("cash[sum]"), 64); err == nil {
      c.Sum = sum
    } else {
      c.Sum = 0.0
    }
  }
}

//
// DEPRECATED
//
func (c *Cash) IsValid() bool {
  c.validateUniquenessOfName()

  return c.Errors().IsEmpty()
}

//
// TODO: test me
//
func (c *Cash) validateUniquenessOfName() {
  var count int

  query := DB.Table("cashes").Where("LOWER(name)=?", strings.ToLower(c.Name))

  if c.ID != 0 {
    query = query.Not("id = ?", c.ID)
  }

  query.Count(&count)

  if count > 0 {
    c.Errors().Add("name", "already exists")
  }
}

//
// CashCollection.Search
//
func (c *CashCollection) Search(values url.Values) {
  DB.Find(c)
}

//
// Cash.Create
//
func (c *Cash) Create() {
  DB.Create(c)
}

//
// Cash.Update
//
func (c *Cash) Update() {
  DB.Save(c)
}
