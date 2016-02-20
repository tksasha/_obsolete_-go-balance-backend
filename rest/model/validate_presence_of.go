package model

import (
  "reflect"

  . "github.com/tksasha/balance/rest/resource"
)

func ValidatePresenceOf(resource Resource, attribute string) {
  reflected := reflect.ValueOf(resource).Elem()

  value := reflect.Indirect(reflected).FieldByName(attribute)

  if value.String() == "" {
    resource.Errors().Add(attribute, "can't be blank")
  }
}
