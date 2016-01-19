package main

type Errors map[string][]string

func (e Errors) Set(key string, value string) {
  e[key] = append(e[key], value)
}

func (e Errors) Get(key string) []string {
  return e[key]
}

func (e Errors) IsEmpty() bool {
  for _, value := range e {
    if len(value) != 0 {
      return false
    }
  }

  return true
}
