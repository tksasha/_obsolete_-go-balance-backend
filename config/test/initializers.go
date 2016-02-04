package config_test

import (
  "net/url"

  _ "github.com/mattn/go-sqlite3"
  "github.com/jinzhu/gorm"

  . "../../config"
  . "../../models"
)

var Values url.Values

func init() {
  DB, _ = gorm.Open("sqlite3", DBWorkDir + "test.sqlite3")

  DB.AutoMigrate(&Item{}, &Category{})

  Values = make(url.Values)
}
