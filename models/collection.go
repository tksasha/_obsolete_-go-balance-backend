package models

import (
  "net/url"
)

type Collection interface {
  Search(url.Values)
}
