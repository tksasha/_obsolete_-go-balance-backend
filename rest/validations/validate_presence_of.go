package validations

import (
  "reflect"
  "time"

  . "github.com/tksasha/balance/rest/resource"
)

const message = "can't be blank"

func ValidatePresenceOf(resource Resource, attribute string) {
  reflected := reflect.ValueOf(resource).Elem()

  value := reflect.Indirect(reflected).FieldByName(attribute)

  switch value.Type().Name() {
  case "string":
    if value.String() == "" {
      resource.Errors().Add(attribute, message)
    }
  case "Time":
    if value.Interface().(time.Time).IsZero() {
      resource.Errors().Add(attribute, message)
    }
  }
}
