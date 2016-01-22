package models

import (
  "github.com/jinzhu/gorm"
)

func VisibleCategories(db *gorm.DB) *gorm.DB {
  return db.Where("visible IN('t', 1)")
}
