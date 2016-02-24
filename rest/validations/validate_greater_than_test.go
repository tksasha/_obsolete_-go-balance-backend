package validations_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/model"
  . "github.com/tksasha/balance/rest/validations"
)

type I struct {
  Model

  SumFloat  float64
  SumInt    int
}

func TestValidateGreaterThanForFloat(t *testing.T) {
  i := new(I)

  ValidateGreaterThan(i, "SumFloat", 0.0)

  assert.Equal(t, []string{"can't be less than or equal `0.000000`"}, i.Errors().Get("SumFloat"))
}

func TestValidateGreaterThanForInt(t *testing.T) {
  i := new(I)

  ValidateGreaterThan(i, "SumInt", 0)

  assert.Equal(t, []string{"can't be less than or equal `0`"}, i.Errors().Get("SumInt"))
}
