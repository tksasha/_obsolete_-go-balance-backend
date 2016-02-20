package handler

import (
  "net/url"

  "github.com/tksasha/balance/rest/errors"
)

type Resource interface {
  Find(string) error
  Build(url.Values)
  Create()
  Update()
  Destroy()
  Errors() *errors.Errors
  IsValid() bool
}
