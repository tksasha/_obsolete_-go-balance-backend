package validations

import (
  "reflect"
  "time"

  . "github.com/tksasha/balance/rest/resource"
)

func ValidatePresenceOf(resource Resource, attribute string) {
  reflected := reflect.ValueOf(resource).Elem()

  value := reflect.Indirect(reflected).FieldByName(attribute)

  switch value.Type().Name() {
  case "string":
    if value.String() == "" {
      resource.Errors().Add(attribute, "can't be blank")
    }
  case "Time":
    if value.Interface().(time.Time).IsZero() {
      resource.Errors().Add(attribute, "can't be blank")
    }
  case "int":
    if value.Int() == 0 {
      resource.Errors().Add(attribute, "can't be blank")
    }
  }
}
