package validations

import (
  "unicode/utf8"
  "reflect"
  "fmt"

  . "github.com/tksasha/balance/rest/resource"
)

func ValidateLengthOf(resource Resource, attribute string, rule string, chars int) {
  value := reflect.ValueOf(resource).Elem()

  value = reflect.Indirect(value).FieldByName(attribute)

  switch rule {
  case "at least":
    if utf8.RuneCountInString(value.String()) < chars {
      message := fmt.Sprintf("length can't be less than %d characters", chars)

      resource.Errors().Add(attribute, message)
    }
  case "at most":
    if utf8.RuneCountInString(value.String()) > chars {
      message := fmt.Sprintf("length can't be more than %d characters", chars)

      resource.Errors().Add(attribute, message)
    }
  case "exactly":
    if utf8.RuneCountInString(value.String()) != chars {
      message := fmt.Sprintf("length must equal %d characters", chars)

      resource.Errors().Add(attribute, message)
    }
  }
}
