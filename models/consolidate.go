package models

import (
  "net/url"

  "github.com/tksasha/go-date"

  . "../config"
)

type Consolidate struct {
  Sum         float32 `json:"sum"`
  CategoryID  int     `json:"category_id"`
}

type ConsolidateCollection []Consolidate

//
// ConsolidateCollection.Search
//
func (c *ConsolidateCollection) Search(values url.Values) {
  d := date.New(values.Get("year"), values.Get("month"))

  DB.
    Table("items").
    Where("date BETWEEN ? AND ?", d.BeginningOfMonth().String(), d.EndOfMonth().String()).
    Select("SUM(sum) AS sum, category_id").
    Group("category_id").
    Scan(c)
}
