package models

type Login struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
