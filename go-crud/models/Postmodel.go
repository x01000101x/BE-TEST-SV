package models

import (
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	Title    string `gorm:"size:200"`
	Content  string `gorm:"type:text"`
	Category string `gorm:"size:100"`
	Status   string `gorm:"size:100"` //Publish | Draft | Thrash
}
