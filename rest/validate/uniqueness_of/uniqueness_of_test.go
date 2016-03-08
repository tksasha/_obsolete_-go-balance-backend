package uniqueness_of_test

import (
  "testing"

  "github.com/stretchr/testify/assert"
  . "github.com/tksasha/balance/rest/validate"
  . "github.com/tksasha/balance/rest/model"
  . "github.com/tksasha/balance/rest/resource"
  . "github.com/tksasha/balance/rest/db"

  _ "github.com/tksasha/balance/config/test"
)

func init() {
  DB.AutoMigrate(&M{})

  DB.Exec("TRUNCATE ms")
}

type M struct {
  Model

  ID    int
  Name  string  `valid:"unique"`
}

func TestValidateUniquenessOfOnCreate(t *testing.T) {
  f := &M { Name: "First" }

  Create(f)

  s := &M { Name: "First" }

  Validate(s)

  assert.Contains(t, s.Errors().Get("Name"), "already exists")
}

func TestValidateUniquenessOfOnUpdate(t *testing.T) {
  m := &M { Name: "Second" }

  Create(m)

  Validate(m)

  assert.NotContains(t, m.Errors().Get("Name"), "already exists")
}
