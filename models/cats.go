package models

import (
	"gorm.io/gorm"
)

type Cats struct {
	gorm.Model
	ID    uint `json:"id" orm:"auto" gorm:"primary_key"`
	Name  string
	Image string
}
