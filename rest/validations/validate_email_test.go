package validations_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/model"
  . "github.com/tksasha/balance/rest/validations"
)

type A struct {
  Model

  Email string
}

func TestValidateEmailWithEmptyValue(t *testing.T) {
  m := new(A)

  ValidateEmail(m)

  assert.Contains(t, m.Errors().Get("Email"), "is invalid")
}

func TestValidateEmailWithInvalidValue(t *testing.T) {
  m := &A{ Email: "invalid" }

  ValidateEmail(m)

  assert.Contains(t, m.Errors().Get("Email"), "is invalid")
}

func TestValidateEmailWithValidValue(t *testing.T) {
  m := &A{ Email: "one@digits.com" }

  ValidateEmail(m)

  assert.NotContains(t, m.Errors().Get("Email"), "is invalid")
}
