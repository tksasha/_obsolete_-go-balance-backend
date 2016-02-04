package config

import (
  "fmt"
  "os"
  "time"

  "github.com/jinzhu/gorm"
)

var WorkDir string

var DBWorkDir string

var DB gorm.DB

func init() {
  //
  // Set TimeZone
  //
  time.Local = time.UTC

  WorkDir = fmt.Sprintf("%s/.balance", os.Getenv("HOME"))

  DBWorkDir = fmt.Sprintf("%s/db", WorkDir)

  //
  // Prepare DBWorkDir
  //
  if err := os.MkdirAll(DBWorkDir, 0755); err != nil {
    panic(err)
  }
}
