package models

import (
  "time"
  "net/url"
  "strconv"

  "github.com/tksasha/go-date"

  . "../config"
)

type Item struct {
  BaseModel

  ID          int       `json:"id"`
  Date        time.Time `json:"date"`
  Sum         float32   `json:"sum"`
  Description string    `json:"description"`
  Category    Category  `json:"category"`
  CategoryID  int       `json:"-"`
}

func (i *Item) Build(values url.Values) {
  i.Init()

  if d := values.Get("item[date]"); d != "" {
    i.Date = date.New(d).Time()
  }

  if category_id := values.Get("item[category_id]"); category_id != "" {
    if id, err := strconv.Atoi(category_id); err == nil {
      i.CategoryID = id
    }
  }

  if sum := values.Get("item[sum]"); sum != "" {
    if s, err := strconv.ParseFloat(sum, 32); err == nil {
      i.Sum = float32(s)
    }
  }
}

func (i *Item) IsValid() bool {
  i.Errors.Add("date", "something went wrong")

  return i.Errors.IsEmpty()
}

func (i *Item) IsCreate(values url.Values) bool {
  i.Build(values)

  if i.IsValid() {
    DB.Create(i)

    return true
  } else {
    return false
  }
}
