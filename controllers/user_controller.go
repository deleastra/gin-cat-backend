package controllers

import (
	"net/http"

	"cat-backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserController represents the controller for operating on the User resource
type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// @Summary Creates a new user
// @Description Creates a new user by parsing a User struct from the request body and saving it to the database.
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "New user details"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /users [post]
func (uc UserController) CreateUser(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Message: "Invalid request payload", Code: http.StatusBadRequest})
		return
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Message: "Error hashing password", Code: http.StatusInternalServerError})
		return
	}
	user.Password = string(hashedPassword)

	// Save the user to the database
	if err := uc.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Message: "Error creating user", Code: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusCreated, models.Response{Message: "Success", Data: user, Code: http.StatusCreated})
}

// @Summary Retrieves a user
// @Description Retrieves a user from the database with the given ID and returns it to the client.
// @Tags users
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [get]
func (uc UserController) GetUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")

	// Find the user with the given ID
	var user models.User
	if err := uc.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{Message: "User not found", Code: http.StatusNotFound})
		return
	}

	// Return the user to the client
	c.JSON(http.StatusOK, models.Response{Message: "Success", Data: user, Code: http.StatusOK})
}

// @Summary Updates a user by ID
// @Description Updates a user with the given ID using the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID of the user to update"
// @Param user body models.User true "Updated user information"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [put]
func (uc UserController) UpdateUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")

	// Find the user with the given ID
	var user models.User
	if err := uc.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{Message: "User not found", Code: http.StatusNotFound})
		return
	}

	// Bind the request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Message: "Invalid request payload", Code: http.StatusBadRequest})
		return
	}

	// Hash the user's password if it was provided in the request body
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Message: "Error hashing password", Code: http.StatusInternalServerError})
			return
		}
		user.Password = string(hashedPassword)
	}

	// Save the updated user to the database
	if err := uc.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Message: "Error updating user", Code: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "Success", Code: http.StatusOK})
}

// @Summary Deletes a user by ID
// @Description Deletes a user with the given ID from the database
// @Tags users
// @Produce json
// @Param id path string true "ID of the user to delete"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /users/{id} [delete]
func (uc UserController) DeleteUser(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")

	// Find the user with the given ID
	var user models.User
	if err := uc.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{Message: "User not found", Code: http.StatusNotFound})
		return
	}

	// Delete the user from the database
	if err := uc.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Message: "Error deleting user", Code: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "Success", Code: http.StatusOK})
}
