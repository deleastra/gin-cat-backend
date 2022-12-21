package controllers

import (
	"net/http"

	"cat-backend/models"

	"github.com/gin-gonic/gin"
)

// GET
func (db *DBController) GetUsers(c *gin.Context) {
	_type := c.Query("type")
	_where := map[string]interface{}{}

	if _type != "" {
		_where["type"] = _type
	}

	var cats []models.Users
	db.Database.Where(_where).Find(&cats)

	c.JSON(http.StatusOK, gin.H{"results": &cats})
}

// GET BY ID
func (db *DBController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user models.Users

	db.Database.First(&user, id)

	c.JSON(http.StatusOK, gin.H{"results": &user})
}

// POST
func (db *DBController) CreateUser(c *gin.Context) {
	var user models.Users
	err := c.ShouldBind(&user)

	result := db.Database.Create(&user)

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &user})
	}
}

// PATCH
func (db *DBController) UpdateUser(c *gin.Context) {

	var user models.Users
	err := c.ShouldBind(&user)

	result := db.Database.Updates(user)

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &user})
	}
}

// DELETE
func (db *DBController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.Users
	db.Database.Delete(&user, id)

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK})
}
