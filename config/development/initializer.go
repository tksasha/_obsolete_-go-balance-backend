package development

import (
  _ "github.com/mattn/go-sqlite3"
  "github.com/jinzhu/gorm"

  . "../../config"
  . "../../models"
)

func init() {
  var err error

  //
  // Make DataBase Connection
  //
  if DB, err = gorm.Open("sqlite3", DBWorkDir + "/development.sqlite3"); err != nil {
    panic(err)
  }

  //
  // Run Migrations
  //
  DB.AutoMigrate(&Item{}, &Category{})

  //
  // Show SQL in LOG
  //
  DB.LogMode(true)
}
