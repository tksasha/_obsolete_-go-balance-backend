package validations

import (
  "reflect"

  "github.com/asaskevich/govalidator"
  . "github.com/tksasha/balance/rest/resource"
)

func ValidateFormatOfEmail(resource Resource) {
  value := reflect.ValueOf(resource).Elem()

  value = reflect.Indirect(value).FieldByName("Email")

  if govalidator.IsEmail(value.String()) == false {
    resource.Errors().Add("Email", "is invalid")
  }
}
