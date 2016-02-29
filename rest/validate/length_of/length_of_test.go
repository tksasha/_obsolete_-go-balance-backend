package validate_length_of_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validate"
  . "github.com/tksasha/balance/rest/model"
)

type M struct {
  Model

  FirstName   string  `valid:"length_at_least(6)"`
  MiddleName  string  `valid:"length_at_most(6)"`
  LastName    string  `valid:"length_exactly(6)"`
}

func TestValidateLengthOfAtLeast(t *testing.T) {
  m := &M{ MiddleName: "1234567", LastName: "1234567" }

  Validate(m)

  assert.Contains(t, m.Errors().Get("FirstName"), "length can't be less than 6 characters")
}

func TestValidateLengthOfAtMost(t *testing.T) {
  m := &M{ MiddleName: "1234567", LastName: "1234567" }

  Validate(m)

  assert.Contains(t, m.Errors().Get("MiddleName"), "length can't be more than 6 characters")
}

func TestValidateLengthOfExactly(t *testing.T) {
  m := &M{ MiddleName: "1234567", LastName: "1234567" }

  Validate(m)

  assert.Contains(t, m.Errors().Get("LastName"), "length must equal 6 characters")
}
