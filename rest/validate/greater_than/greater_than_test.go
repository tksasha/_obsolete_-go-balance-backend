package greater_than_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validate"
  . "github.com/tksasha/balance/rest/model"
)

type M struct {
  Model

  Age   int     `valid:"greater_than(18)"`
  Sum   float64 `valid:"greater_than(0)"`
  Rate  float64 `valid:"greater_than(9.99)"`
}

func TestValidateGreaterThanForInt(t *testing.T) {
  m := new(M)

  Validate(m)

  assert.Contains(t, m.Errors().Get("Age"), "can't be less than or equal `18`")
}

func TestValidateGreaterThanForFloat(t *testing.T) {
  m := new(M)

  Validate(m)

  assert.Contains(t, m.Errors().Get("Sum"), "can't be less than or equal `0.000000`")

  assert.Contains(t, m.Errors().Get("Rate"), "can't be less than or equal `9.990000`")
}
