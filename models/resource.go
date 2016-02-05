package models

type Resource interface {
  First()
}

func First(r Resource) {
  r.First()
}
