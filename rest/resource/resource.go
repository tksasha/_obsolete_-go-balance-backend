package resource

import (
  "net/url"

  "github.com/tksasha/balance/rest/errors"
  . "github.com/tksasha/balance/rest/db"
)

type Resource interface {
  Build(url.Values)
  Create()
  Update()
  Delete()
  Errors() *errors.Errors
}

func Find(resource Resource, id string) error {
  return DB.First(resource, id).Error
}
