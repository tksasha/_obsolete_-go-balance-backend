package validations_test

import (
  "time"

  . "github.com/tksasha/balance/rest/model"
)

type M struct {
  Model

  FullName  string
  Date      time.Time
  Age       int
  Email     string
  SumFloat  float64
  SumInt    int
}
