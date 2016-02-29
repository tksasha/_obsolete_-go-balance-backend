package greater_than_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validate"
  . "github.com/tksasha/balance/rest/model"
)

type M struct {
  Model

  Age int `valid:"greater_than(0)"`
}

func TestValidateGreaterThan(t *testing.T) {
  m := new(M)

  Validate(m)

  assert.Contains(t, m.Errors().Get("Age"), "can't be less than or equal `0`")
}
