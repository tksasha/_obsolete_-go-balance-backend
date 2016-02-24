package validations

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
    value := toFloat64(value)

    if v.Interface().(float64) <= value {
      message := fmt.Sprintf("can't be less than or equal `%f`", value)

      resource.Errors().Add(attribute, message)
    }
  case int:
    value := toInt(value)

    if v.Interface().(int) <= value {
      message := fmt.Sprintf("can't be less than or equal `%d`", value)

      resource.Errors().Add(attribute, message)
    }
  }
}

func toFloat64(i interface{}) float64 {
  value := reflect.ValueOf(i)

  value = reflect.Indirect(value)

  floatType := reflect.TypeOf(float64(0))

  if value.Type().ConvertibleTo(floatType) {
    return value.Convert(floatType).Float()
  } else {
    return 0.0
  }
}

func toInt(i interface{}) int {
  value := reflect.ValueOf(i)

  value = reflect.Indirect(value)

  intType := reflect.TypeOf(int(0))

  if value.Type().ConvertibleTo(intType) {
    return int(value.Convert(intType).Int())
  } else {
    return 0
  }
}
