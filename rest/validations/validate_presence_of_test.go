package validations_test

import (
  "testing"
  "time"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/balance/rest/validations"
  . "github.com/tksasha/balance/rest/model"
)

type M struct {
  Model

  Name      string
  Date      time.Time
  Age       int
}

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
