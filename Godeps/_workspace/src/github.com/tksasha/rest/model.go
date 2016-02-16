package rest

import (
	"github.com/tksasha/balance/Godeps/_workspace/src/github.com/tksasha/rest/errors"
)

type Model struct {
	Errors errors.Errors `sql:"-" json:"-"` // Ignore this field
}

func (model *Model) GetErrors() errors.Errors {
	return model.Errors
}

func (m *Model) Find(id string) error { return nil }

func (m *Model) Destroy() {}
