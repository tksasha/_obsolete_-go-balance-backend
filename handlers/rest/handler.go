package rest_handler

import (
  "net/http"
  "net/url"
  "encoding/json"

  "github.com/julienschmidt/httprouter"
  "github.com/tksasha/go-errors"
)

type Resource interface {
  IsCreate(url.Values) bool
  IsUpdate(url.Values) bool
  GetErrors() errors.Errors
  Find(string) error
  Destroy()
}

type Collection interface {
  Search(url.Values)
}

type RESTHandler struct {
  Resource    func() Resource
  Collection  func() Collection
}

//
// GET - Show Resource
//

//
// GET - List Collection
//
func (h RESTHandler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  collection := h.Collection()

  collection.Search(r.URL.Query())

  render(w, collection, 200)
}

//
// POST - Create Resource
//
func (h RESTHandler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  resource := h.Resource()

  if resource.IsCreate(values(r, params)) {
    render(w, resource, 200)
  } else {
    render(w, resource.GetErrors(), 422)
  }
}

//
// DELETE - Destroy Resource
//
func (h RESTHandler) Destroy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  resource := h.Resource()

  if err := resource.Find(params.ByName("id")); err != nil && err.Error() == "record not found" {
    http.NotFound(w, r)

    return
  }

  resource.Destroy()

  render(w, "OK", 200)
}

//
// PATCH - Update Resource
//
func(h RESTHandler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  resource := h.Resource()

  if err := resource.Find(params.ByName("id")); err != nil && err.Error() == "record not found" {
    http.NotFound(w, r)

    return
  }

  if resource.IsUpdate(values(r, params)) {
    render(w, resource, 200)
  } else {
    render(w, resource.GetErrors(), 422)
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
