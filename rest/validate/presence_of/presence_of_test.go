package presence_of_test

import (
  "testing"
  "time"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/balance/rest/validate"
  . "github.com/tksasha/balance/rest/model"
)

type M struct {
  Model

  FullName  string    `valid:"present"`
  Date      time.Time `valid:"present"`
  Age       int       `valid:"present"`
}

func TestValidatePresenceOfString(t *testing.T) {
  m := new(M)

  Validate(m)

  assert.Contains(t, m.Errors().Get("FullName"), "can't be blank")
}

func TestValidatePresenceOfTime(t *testing.T) {
  m := new(M)

  Validate(m)

  assert.Contains(t, m.Errors().Get("Date"), "can't be blank")
}

func TestValidatePresenceOfInt(t *testing.T) {
  m := new(M)

  Validate(m)

  assert.Contains(t, m.Errors().Get("Age"), "can't be blank")
}
