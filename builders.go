package main

import (
  "net/url"
  "strconv"

  "github.com/jinzhu/now"
)

//
// Category
//
func (c *Category) Build(values url.Values) {
  c.Name = values.Get("category[name]")
}

//
// Item
//
func (i *Item) Build(values url.Values) {
  if t, err := now.Parse(values.Get("item[date]")); err == nil {
    i.Date = t
  }

  if id, err := strconv.Atoi(values.Get("item[category_id]")); err == nil {
    i.CategoryID = id
  }

  if sum, err := strconv.ParseFloat(values.Get("item[sum]"), 32); err == nil {
    i.Sum = float32(sum)
  }
}
