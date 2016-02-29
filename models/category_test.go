package models_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validate"
  . "github.com/tksasha/balance/models"
)

func TestCategoryPresenceOfName(t *testing.T) {
  c := new(Category)

  Validate(c)

  assert.Contains(t, c.Errors().Get("Name"), "can't be blank")
}
