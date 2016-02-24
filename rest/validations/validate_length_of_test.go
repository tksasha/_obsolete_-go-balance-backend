package validations_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validations"
)

func TestValidateLengthOfWithMinimum(t *testing.T) {
  m := new(M)

  ValidateLengthOf(m, "Name", "at least", 6)

  assert.Contains(t, m.Errors().Get("Name"), "length can't be less than 6 characters")
}

func TestValidateLengthOfWithMaximum(t *testing.T) {
  m := &M{ Name: "1234567" }

  ValidateLengthOf(m, "Name", "at most", 6)

  assert.Contains(t, m.Errors().Get("Name"), "length can't be more than 6 characters")
}

func TestValidateLengthOfWithExactlyValue(t *testing.T) {
  m := &M{ Name: "1234567" }

  ValidateLengthOf(m, "Name", "exactly", 6)

  assert.Contains(t, m.Errors().Get("Name"), "length must equal 6 characters")
}
