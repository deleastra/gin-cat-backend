package models

import (
	"gorm.io/gorm"
)

// Cats represents a type of data object for storing information about cats.
type Cats struct {
	gorm.Model
	ID    uint   `gorm:"primary_key;auto_increment" json:"id"` // ID is the primary key and auto-incrementing field for the cats table in the database.
	Name  string `json:"name"`                                 // Name is the name of the cat.
	Image string `json:"image"`                                // Image is the file path or URL of an image for the cat.
}
