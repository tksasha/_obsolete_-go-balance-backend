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

type ItemCollection []Item

//
// Item.Build()
//
func (item *Item) Build(values url.Values) {
  item.Init()

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
// Item.IsValid()
//
func (i *Item) IsValid() bool {
  i.validateSumGreaterThanZero()

  i.validatePresenceOfCategory()

  i.validatePresenceOfDate()

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
// Item.IsUpdate
//
func (i *Item) IsUpdate(values url.Values) bool {
  i.Build(values)

  if i.IsValid() {
    DB.Save(i)

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

func (i *Item) validateSumGreaterThanZero() {
  if i.Sum == 0.0 {
    i.Errors.Add("sum", "must be greater than 0")
  }
}

func (i *Item) validatePresenceOfCategory() {
  var category Category

  if DB.First(&category, i.CategoryID).RecordNotFound() {
    i.Errors.Add("category_id", "can't be blank")
  }
}

func (i *Item) validatePresenceOfDate() {
  if i.Date.IsZero() {
    i.Errors.Add("date", "can't be blank")
  }
}

func (i *Item) First() {
  DB.First(&i)
}

func (i *ItemCollection) Search(values url.Values) {
  d := date.New(values.Get("year"), values.Get("month"))

  DB.
    Where("date BETWEEN ? AND ?", d.BeginningOfMonth().String(), d.EndOfMonth().String()).
    Order("date").
    Preload("Category").
    Find(&i)
}
