package errors

import (
  "encoding/json"

  "github.com/serenize/snaker"
)

type Errors struct {
  collection map[string][]string
}

func (e *Errors) Add(key string, value string) {
  if len(e.collection) == 0 {
    e.collection = make(map[string][]string)
  }

  e.collection[key] = append(e.collection[key], value)
}

func (e *Errors) Get(key string) []string {
  return e.collection[key]
}

func (e *Errors) IsEmpty() bool {
  for _, value := range e.collection {
    if len(value) != 0 {
      return false
    }
  }

  return true
}

func (e Errors) MarshalJSON() ([]byte, error) {
  for key, value := range e.collection {
    delete(e.collection, key)

    e.collection[snaker.CamelToSnake(key)] = value
  }

  return json.Marshal(e.collection)
}
