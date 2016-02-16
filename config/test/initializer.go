package test

import (
	"net/url"
)

var Values url.Values

func init() {
	Values = make(url.Values)
}
