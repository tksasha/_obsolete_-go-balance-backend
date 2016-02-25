package validations_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validations"
)

func TestValidateLengthOfWithMinimum(t *testing.T) {
  m := new(M)

  ValidateLengthOf(m, "FullName", "at least", 6)

  assert.Contains(t, m.Errors().Get("FullName"), "length can't be less than 6 characters")
}

func TestValidateLengthOfWithMaximum(t *testing.T) {
  m := &M{ FullName: "1234567" }

  ValidateLengthOf(m, "FullName", "at most", 6)

  assert.Contains(t, m.Errors().Get("FullName"), "length can't be more than 6 characters")
}

func TestValidateLengthOfWithExactlyValue(t *testing.T) {
  m := &M{ FullName: "1234567" }

  ValidateLengthOf(m, "FullName", "exactly", 6)

  assert.Contains(t, m.Errors().Get("FullName"), "length must equal 6 characters")
}
