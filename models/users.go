package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model          // Embeds fields from the gorm.Model struct, including ID, CreatedAt, and UpdatedAt
	Username    *string `gorm:"unique;size:12"` // Unique username with a size of 12 characters
	Password    *string `json:"-"`              // Password for the user (not serialized as part of a JSON response)
	DisplayName *string `gorm:"size:24"`        // Display name for the user with a size of 24 characters
}
