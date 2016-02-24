package validations_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validations"
)

func TestValidateEmailWithEmptyValue(t *testing.T) {
  m := new(M)

  ValidateFormatOfEmail(m)

  assert.Contains(t, m.Errors().Get("Email"), "is invalid")
}

func TestValidateEmailWithInvalidValue(t *testing.T) {
  m := &M{ Email: "invalid" }

  ValidateFormatOfEmail(m)

  assert.Contains(t, m.Errors().Get("Email"), "is invalid")
}

func TestValidateEmailWithValidValue(t *testing.T) {
  m := &M{ Email: "one@digits.com" }

  ValidateFormatOfEmail(m)

  assert.NotContains(t, m.Errors().Get("Email"), "is invalid")
}
