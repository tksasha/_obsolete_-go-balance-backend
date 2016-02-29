package validations

import (
  "reflect"

  "github.com/asaskevich/govalidator"
  . "github.com/tksasha/balance/rest/resource"
)

func ValidateFormatOfEmail(resource Resource, attribute string) {
  value := reflect.ValueOf(resource).Elem()

  value = reflect.Indirect(value).FieldByName(attribute)

  if govalidator.IsEmail(value.String()) == false {
    resource.Errors().Add(attribute, "is not email")
  }
}
