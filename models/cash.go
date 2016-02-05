package models

import (
  "net/url"
  "strconv"
  "strings"

  . "../config"
)

type Cash struct {
  BaseModel

  ID    int
  Name  string
  Sum   float64
}

type CashCollection []Cash

//
// Cash.Build
//
func (c *Cash) Build(values url.Values) {
  c.Init()

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
// Cash.IsValid
//
func (c *Cash) IsValid() bool {
  c.validatePresenceOfName()

  c.validateUniquenessOfName()

  return c.Errors.IsEmpty()
}

//
// Cash.IsCreate
//
func (c *Cash) IsCreate(values url.Values) bool {
  c.Build(values)

  if c.IsValid() {
    DB.Create(c)

    return true
  } else {
    return false
  }
}

//
// Cash.IsUpdate
//
func (c *Cash) IsUpdate(values url.Values) bool {
  return true
}

func (c *Cash) validatePresenceOfName() {
  if c.Name == "" {
    c.Errors.Add("name", "can't be blank")
  }
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
    c.Errors.Add("name", "already exists")
  }
}

//
// CashCollection.Search
//
func (c *CashCollection) Search(values url.Values) {
  DB.Find(c)
}
