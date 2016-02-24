package validations_test

import (
  "testing"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/balance/rest/validations"
)

func TestValidatePresenceOfString(t *testing.T) {
  m := new(M)

  ValidatePresenceOf(m, "Name")

  assert.Equal(t, []string{"can't be blank"}, m.Errors().Get("Name"))
}

func TestValidatePresenceOfTime(t *testing.T) {
  m := new(M)

  ValidatePresenceOf(m, "Date")

  assert.Equal(t, []string{"can't be blank"}, m.Errors().Get("Date"))
}

func TestValidatePresenceOfInt(t *testing.T) {
  m := new(M)

  ValidatePresenceOf(m, "Age")

  assert.Equal(t, []string{"can't be blank"}, m.Errors().Get("Age"))
}
