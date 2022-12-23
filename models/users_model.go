package models

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model         // Embeds fields from the gorm.Model struct, including ID, CreatedAt, and UpdatedAt
	Email       string `gorm:"not null" json:"email"`
	Username    string `gorm:"not null" json:"username"`
	Password    string `gorm:"not null" json:"password"`
	DisplayName string `gorm:"not null" json:"display_name"`
}

// CheckPassword checks if the provided password matches the user's hashed password.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// GenerateJWT generates a JWT token for the user.
func (u *User) GenerateJWT() (string, error) {
	// Set the claims for the JWT token.
	claims := jwt.MapClaims{
		"id":           u.ID,
		"email":        u.Email,
		"display_name": u.DisplayName,
		"exp":          time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create the JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key.
	secretKey := []byte("my-secret-key")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
