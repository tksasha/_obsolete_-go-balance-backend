package main

import (
	"time"
)

type Item struct {
  ID          int       `json:"id"`
  Date        time.Time `json:"date"`
  Sum         float32   `json:"sum"`
  Formula     string    `json:"formula"`
  Description string    `json:"description"`
}

type Consolidate struct {
  Sum         float32 `json:"sum"`
  CategoryID  int     `json:"category_id"`
}

type Category struct {
  ID      int     `json:"id"`
  Name    string  `json:"name"`
  Income  bool    `json:"income"`
  Visible bool    `json:"visible" sql:"default:'t'"`
}
