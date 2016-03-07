package resource

import (
  "net/url"

  "github.com/tksasha/balance/rest/errors"
  . "github.com/tksasha/balance/rest/db"
)

type Resource interface {
  Build(url.Values)
  Errors() *errors.Errors
}

func Find(resource Resource, id string) error {
  return DB.First(resource, id).Error
}

func Create(resource Resource) {
  DB.Create(resource)
}

func Update(resource Resource) {
  DB.Save(resource)
}

func Delete(resource Resource) {
  DB.Delete(resource)
}
