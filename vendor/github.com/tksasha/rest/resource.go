package rest

import (
  "net/url"

  "github.com/tksasha/rest/errors"
)

type Resource interface {
  IsCreate(url.Values) bool
  IsUpdate(url.Values) bool
  GetErrors() errors.Errors
  Find(string) error
  Destroy()
}
