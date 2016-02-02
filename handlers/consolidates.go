package handlers

import (
  "net/http"

  "github.com/julienschmidt/httprouter"

  . "../config"
  . "../models"
)

type Consolidates struct {
  BaseHandler
}

func (c Consolidates) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  var consolidates []Consolidate

  DB.
    Table("items").
    Select("SUM(sum) AS sum, category_id").
    Group("category_id").
    Scan(&consolidates)

  c.render(w, consolidates, 200)
}
