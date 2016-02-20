package model

import (
  "net/url"

  "github.com/tksasha/balance/rest/errors"
)

type Model struct {
  errors errors.Errors  `sql:"-" json:"-"` // Ignore this field
}

func (m *Model) Errors() *errors.Errors { return &m.errors }

func (*Model) Find(id string) error { return nil }

func (*Model) Build(url.Values) {}

func (*Model) Create() {}

func (*Model) Update() {}

func (*Model) Destroy() {}

func (*Model) IsValid() bool { return true }
