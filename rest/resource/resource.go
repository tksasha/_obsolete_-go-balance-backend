package resource

import (
  "net/url"

  "github.com/tksasha/balance/rest/errors"
)

type Resource interface {
  Find(string) error
  Build(url.Values)
  Create()
  Update()
  Delete()
  Errors() *errors.Errors
}
