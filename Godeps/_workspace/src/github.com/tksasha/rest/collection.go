package rest

import (
  "net/url"
)

type Collection interface {
  Search(url.Values)
}
