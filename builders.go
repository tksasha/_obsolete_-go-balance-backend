package main

import (
  "net/url"
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
}
