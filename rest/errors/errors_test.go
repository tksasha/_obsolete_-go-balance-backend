package errors_test

import (
  "testing"
  "encoding/json"

  "github.com/stretchr/testify/assert"

  . "github.com/tksasha/balance/rest/errors"
)

func TestAdd(t *testing.T) {
  var e Errors

  e.Add("name", "can't be blank")

  assert.Equal(t, []string{ "can't be blank" }, e.Get("name"))
}

func TestGetWithUnexistedKey(t *testing.T) {
  var e Errors

  assert.Equal(t, []string(nil), e.Get("foo"))
}

func TestIsEmpty(t *testing.T) {
  var e Errors

  assert.True(t, e.IsEmpty())

  e.Add("name", "can't be blank")

  assert.False(t, e.IsEmpty())
}

func TestMarshalJSON(t *testing.T) {
  var e Errors

  e.Add("name", "can't be blank")

  marshaled, _ := json.Marshal(e)

  assert.Equal(t, string(marshaled), `{"name":["can't be blank"]}`)
}
