package uniqueness_of

import (
  "reflect"
  "strings"

  "github.com/serenize/snaker"
  . "github.com/tksasha/balance/rest/resource"
  . "github.com/tksasha/balance/rest/db"
)

func ValidateUniquenessOf(resource Resource, attribute string) {
  reflected := reflect.ValueOf(resource).Elem()

  value := reflect.Indirect(reflected).FieldByName(attribute)

  field := snaker.CamelToSnake(attribute)

  var count int

  query := DB.Model(resource).Where("LOWER(" + field + ")=?", strings.ToLower(value.String()))

  id := reflect.Indirect(reflected).FieldByName("ID").Int()

  if id > 0 {
    query = query.Not("id=?", id)
  }

  query.Count(&count)

  if count > 0 {
    resource.Errors().Add(attribute, "already exists")
  }
}
