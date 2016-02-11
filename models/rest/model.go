package rest_model

import (
  "github.com/tksasha/go-errors"
)

type RESTModel struct {
  Errors errors.Errors  `sql:"-" json:"-"` // Ignore this field
}

func (model *RESTModel) GetErrors() errors.Errors {
  return model.Errors
}

func (m *RESTModel) Find(id string) error {
  return nil
}

func (m *RESTModel) Destroy() {}
