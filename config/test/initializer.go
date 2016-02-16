package test

import (
	"net/url"

	"github.com/tksasha/balance/Godeps/_workspace/src/github.com/jinzhu/gorm"
	_ "github.com/tksasha/balance/Godeps/_workspace/src/github.com/mattn/go-sqlite3"

	. "github.com/tksasha/balance/config"
	. "github.com/tksasha/balance/models"
)

var Values url.Values

func init() {
	DB, _ = gorm.Open("sqlite3", DBWorkDir+"test.sqlite3")

	DB.AutoMigrate(&Item{}, &Category{})

	Values = make(url.Values)
}
