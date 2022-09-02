package controllers

import (
	"net/http"

	"example.com/golang-restfulapi/models"
	"github.com/gin-gonic/gin"
)

// GET
func (db *DBController) GetCats(c *gin.Context) {
	_type := c.Query("type")
	_where := map[string]interface{}{}

	if _type != "" {
		_where["type"] = _type
	}

	var cats []models.Cats
	db.Database.Where(_where).Find(&cats)

	c.JSON(http.StatusOK, gin.H{"results": &cats})
}

// GET BY ID
func (db *DBController) GetCatByID(c *gin.Context) {
	id := c.Param("id")
	var cats models.Cats

	db.Database.First(&cats, id)
	c.JSON(http.StatusOK, gin.H{"results": &cats})
}

// POST
func (db *DBController) CreateCat(c *gin.Context) {
	var cats models.Cats
	err := c.ShouldBind(&cats)

	result := db.Database.Create(&cats)

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &cats})
	}
}

// PATCH
func (db *DBController) UpdateCat(c *gin.Context) {

	var cat models.Cats
	err := c.ShouldBind(&cat)

	result := db.Database.Updates(cat)

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &cat})
	}
}

// DELETE
func (db *DBController) DeleteCat(c *gin.Context) {
	id := c.Param("id")
	var cats models.Cats
	db.Database.Delete(&cats, id)

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK})
}
