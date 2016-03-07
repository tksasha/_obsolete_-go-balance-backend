package handler

import (
  "net/http"
  "net/url"
  "encoding/json"

  "github.com/julienschmidt/httprouter"

  . "github.com/tksasha/balance/rest/resource"
  . "github.com/tksasha/balance/rest/collection"
  . "github.com/tksasha/balance/rest/validate"
)

type Handler struct {
  Resource    func() Resource
  Collection  func() Collection
}

//
// GET - Show Resource
//

//
// GET - List Collection
//
func (h Handler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  collection := h.Collection()

  collection.Search(r.URL.Query())

  render(w, collection, 200)
}

//
// POST - Create Resource
//
func (h Handler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  resource := h.Resource()

  resource.Build(values(r, params))

  Validate(resource)

  if resource.Errors().IsEmpty() {
    Create(resource)

    render(w, resource, 200)
  } else {
    render(w, resource.Errors(), 422)
  }
}

//
// DELETE - Destroy Resource
//
func (h Handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  resource := h.Resource()

  err := Find(resource, params.ByName("id"))

  if err != nil && err.Error() == "record not found" {
    http.NotFound(w, r)

    return
  }

  Delete(resource)

  render(w, "OK", 200)
}

//
// PATCH - Update Resource
//
func(h Handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  resource := h.Resource()

  err := Find(resource, params.ByName("id"))

  if err != nil && err.Error() == "record not found" {
    http.NotFound(w, r)

    return
  }

  resource.Build(values(r, params))

  Validate(resource)

  if resource.Errors().IsEmpty() {
    resource.Update()

    render(w, resource, 200)
  } else {
    render(w, resource.Errors(), 422)
  }
}

//
// Service Methods
//
func render(w http.ResponseWriter, items interface{}, code int) {
  w.WriteHeader(code)

  if err := json.NewEncoder(w).Encode(items); err != nil {
    panic(err)
  }
}

func values(r *http.Request, params httprouter.Params) url.Values {
  r.ParseForm()

  values := r.Form

  for _, param := range params {
    values.Set(param.Key, param.Value)
  }

  return values
}
