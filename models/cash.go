package models

import (
  "net/url"
  "strconv"

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
  Name string   `valid:"present,unique"`
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
// CashCollection.Search
//
func (c *CashCollection) Search(values url.Values) {
  DB.Find(c)
}
