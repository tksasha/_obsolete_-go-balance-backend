package models

import (
  "github.com/tksasha/go-errors"
)

type BaseModel struct {
  Errors errors.Errors  `sql:"-" json:"-"` // Ignore this field
}

func (model *BaseModel) Init() {
  model.Errors = make(errors.Errors)
}

func (model *BaseModel) GetErrors() errors.Errors {
  return model.Errors
}
