package models

import (
  "time"
  "net/url"
  "strconv"
  "encoding/json"

  "github.com/tksasha/go-date"

  . "../config"
)

type Item struct {
  BaseModel

  ID          int
  Date        time.Time
  Sum         float64
  Description string
  Category    Category
  CategoryID  int
}

type ItemDecorator struct {
  ID          int       `json:"id"`
  Date        string    `json:"date"`
  Sum         float64   `json:"sum"`
  Description string    `json:"description"`
  Category    Category  `json:"category"`
}

//
// Item.Build()
//
func (item *Item) Build(values url.Values) {
  item.Init()

  if d := values.Get("item[date]"); d != "" {
    item.Date = date.New(d).Time()
  }

  if category_id := values.Get("item[category_id]"); category_id != "" {
    if id, err := strconv.Atoi(category_id); err == nil {
      item.CategoryID = id
    }
  }

  if sum := values.Get("item[sum]"); sum != "" {
    if s, err := strconv.ParseFloat(sum, 64); err == nil {
      item.Sum = s
    }
  }
}

//
// Item.IsValid()
//
func (i *Item) IsValid() bool {
  i.Errors.Add("date", "something went wrong")

  return i.Errors.IsEmpty()
}

//
// Item.IsCreate()
//
func (i *Item) IsCreate(values url.Values) bool {
  i.Build(values)

  if i.IsValid() {
    DB.Create(i)

    return true
  } else {
    return false
  }
}

//
// Item.MarshalJSON()
//
func (i Item) MarshalJSON() ([]byte, error) {
  d := ItemDecorator {
    ID:           i.ID,
    Date:         i.Date.Format("2006-01-02"),
    Sum:          i.Sum,
    Description:  i.Description,
    Category:     i.Category,
  }

  return json.Marshal(d)
}
