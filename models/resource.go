package models

import (
  "net/url"

  "github.com/tksasha/go-errors"
)

type Resource interface {
  IsCreate(url.Values) bool
  IsUpdate(url.Values) bool
  GetErrors() errors.Errors
}
