package handler_test

import (
  "testing"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/balance/rest/handler"
)

func TestNothing(t *testing.T) {
  assert.NotPanics(t, func() {
    _ = &Handler{}
  })
}
