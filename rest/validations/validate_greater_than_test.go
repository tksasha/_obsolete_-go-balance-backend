package validations_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validations"
)

func TestValidateGreaterThanForFloat(t *testing.T) {
  m := new(M)

  ValidateGreaterThan(m, "SumFloat", 0.0)

  assert.Contains(t, m.Errors().Get("SumFloat"), "can't be less than or equal `0.000000`")
}

func TestValidateGreaterThanForInt(t *testing.T) {
  m := new(M)

  ValidateGreaterThan(m, "SumInt", 0)

  assert.Contains(t, m.Errors().Get("SumInt"), "can't be less than or equal `0`")
}
