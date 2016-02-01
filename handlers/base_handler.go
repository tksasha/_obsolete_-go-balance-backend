package handlers

import (
  "net/http"
  "encoding/json"
  "log"
)

const (
  StatusUnprocessableEntity = 422
)

type BaseHandler int

func (BaseHandler) render(w http.ResponseWriter, items interface{}, code ...int) {
  if len(code) == 0 {
    w.WriteHeader(http.StatusOK)
  } else {
    w.WriteHeader(code[0])
  }

  if err := json.NewEncoder(w).Encode(items); err != nil {
    log.Fatalln(err)
  }
}
