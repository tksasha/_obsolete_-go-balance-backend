package models_test

import (
  "testing"
  "encoding/json"

  "github.com/stretchr/testify/assert"
  "github.com/tksasha/go-date"

  . "../models"
)

var item Item

var category Category

func init() {
  category = Category {
    ID: 36,
    Name: "Food",
  }

  item = Item {
    ID: 57,
    Date: date.New("1982-05-17").Time(),
    Sum: 10.52,
    Description: "Red Dry Wine",
    Category: category,
  }
}

func TestMarshalJSON(t *testing.T) {
  assert := assert.New(t)

  m, _ := json.Marshal(item)

  var d ItemDecorator

  if err := json.Unmarshal(m, &d); err != nil {
    panic(err)
  }

  assert.Equal(57, d.ID)

  assert.Equal("1982-05-17", d.Date)

  assert.Equal(10.52, d.Sum)

  assert.Equal("Red Dry Wine", d.Description)

  assert.Equal(36, d.Category.ID)

  assert.Equal("Food", d.Category.Name)
}
