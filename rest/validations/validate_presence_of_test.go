package validations_test

import (
  "testing"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/balance/rest/validations"
)

func TestValidatePresenceOfString(t *testing.T) {
  m := new(M)

  ValidatePresenceOf(m, "FullName")

  assert.Contains(t, m.Errors().Get("FullName"), "can't be blank")
}

func TestValidatePresenceOfTime(t *testing.T) {
  m := new(M)

  ValidatePresenceOf(m, "Date")

  assert.Contains(t, m.Errors().Get("Date"), "can't be blank")
}

func TestValidatePresenceOfInt(t *testing.T) {
  m := new(M)

  ValidatePresenceOf(m, "Age")

  assert.Contains(t, m.Errors().Get("Age"), "can't be blank")
}
