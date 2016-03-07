package models

import (
  "encoding/json"
  "net/url"
  "strconv"
  "time"

  "github.com/tksasha/date"

  . "github.com/tksasha/balance/rest/model"
  . "github.com/tksasha/balance/rest/db"
  _ "github.com/tksasha/balance/config"
)

func init() {
  DB.AutoMigrate(&Item{})
}

type Item struct {
  Model

  ID          int
  Date        time.Time `valid:"present"`
  Sum         float64   `valid:"greater_than(0)"`
  Description string
  Category    Category
  CategoryID  int
}

type ItemDecorator struct {
  ID          int      `json:"id"`
  Date        string   `json:"date"`
  Sum         float64  `json:"sum"`
  Description string   `json:"description"`
  Category    Category `json:"category"`
}

type ItemCollection []Item

//
// Item.Build()
//
func (item *Item) Build(values url.Values) {
  if len(values["item[date]"]) != 0 {
    if value := values.Get("item[date]"); value == "" {
      item.Date = time.Time{}
    } else {
      item.Date = date.New(value).Time()
    }
  }

  if len(values["item[category_id]"]) != 0 {
    if id, err := strconv.Atoi(values.Get("item[category_id]")); err == nil {
      item.CategoryID = id
    } else {
      item.CategoryID = 0
    }
  }

  if len(values["item[sum]"]) != 0 {
    if sum, err := strconv.ParseFloat(values.Get("item[sum]"), 64); err == nil {
      item.Sum = sum
    } else {
      item.Sum = 0.0
    }
  }

  if len(values["item[description]"]) != 0 {
    item.Description = values.Get("item[description]")
  }
}

//
// Item.IsValid() [DEPRECATED]
//
func (i *Item) IsValid() bool {
  i.validatePresenceOfCategory()

  return i.Errors().IsEmpty()
}

//
// Item.MarshalJSON()
//
func (i Item) MarshalJSON() ([]byte, error) {
  d := ItemDecorator{
    ID:          i.ID,
    Date:        i.Date.Format("2006-01-02"),
    Sum:         i.Sum,
    Description: i.Description,
    Category:    i.Category,
  }

  return json.Marshal(d)
}

func (i *Item) validatePresenceOfCategory() {
  var category Category

  if DB.First(&category, i.CategoryID).RecordNotFound() {
    i.Errors().Add("category_id", "can't be blank")
  }
}

//
// ItemCollection.Search
//
func (i *ItemCollection) Search(values url.Values) {
  d := date.New(values.Get("year"), values.Get("month"))

  DB.
    Where("date BETWEEN ? AND ?", d.BeginningOfMonth().String(), d.EndOfMonth().String()).
    Order("date").
    Preload("Category").
    Find(&i)
}
