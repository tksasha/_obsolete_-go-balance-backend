package validations

import (
  "reflect"
  "fmt"

  . "github.com/tksasha/balance/rest/resource"
)

func ValidateGreaterThan(resource Resource, attribute string, value float64) {
  reflected := reflect.ValueOf(resource).Elem()

  if reflect.Indirect(reflected).FieldByName(attribute).Float() <= value {
    message := fmt.Sprintf("can't be less or equal `%f`", value)

    resource.Errors().Add(attribute, message)
  }
}
