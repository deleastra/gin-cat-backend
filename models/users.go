package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID          uint   `json:"id" orm:"auto" gorm:"primary_key"`
	Username    string `gorm:"unique;size:12"`
	password    string
	DisplayName string `gorm:"size:24"`
}
