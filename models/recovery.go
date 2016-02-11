package models

import (
  "net/url"
  "encoding/json"

  . "../config"
)

type Recovery struct {
  BaseModel
}

func (r *Recovery) IsCreate(values url.Values) bool {
  var category Category

  if DB.Unscoped().First(&category, values["category_id"]).RecordNotFound() {
    r.Errors.Add("category_id", "can't be blank")

    return false
  } else {
    category.DeletedAt = nil

    DB.Unscoped().Save(&category)

    return true
  }
}

func (r *Recovery) IsUpdate(values url.Values) bool {
  return true
}

func (*Recovery) MarshalJSON() ([]byte, error) {
  return json.Marshal("OK")
}
