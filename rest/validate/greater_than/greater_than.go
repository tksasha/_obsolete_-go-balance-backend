package greater_than

import (
  "reflect"
  "fmt"

  . "github.com/tksasha/balance/rest/resource"
)

func ValidateGreaterThan(resource Resource, attribute string, value interface{}) {
  v := reflect.ValueOf(resource).Elem()

  v = reflect.Indirect(v).FieldByName(attribute)

  switch value.(type) {
  case float64:
    if v.Interface().(float64) <= value.(float64) {
      message := fmt.Sprintf("can't be less than or equal `%f`", value)

      resource.Errors().Add(attribute, message)
    }
  case int:
    if v.Interface().(int) <= value.(int) {
      message := fmt.Sprintf("can't be less than or equal `%d`", value)

      resource.Errors().Add(attribute, message)
    }
  }
}
