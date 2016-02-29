package format_of_email_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validate"
  . "github.com/tksasha/balance/rest/model"
)

type M struct {
  Model

  Email string `valid:"email"`
}

func TestValidateEmailWithEmptyValue(t *testing.T) {
  m := new(M)

  Validate(m)

  assert.Contains(t, m.Errors().Get("Email"), "is not email")
}

func TestValidateEmailWithInvalidValue(t *testing.T) {
  m := &M{ Email: "invalid" }

  Validate(m)

  assert.Contains(t, m.Errors().Get("Email"), "is not email")
}

func TestValidateEmailWithValidValue(t *testing.T) {
  m := &M{ Email: "one@digits.com" }

  Validate(m)

  assert.NotContains(t, m.Errors().Get("Email"), "is not email")
}
