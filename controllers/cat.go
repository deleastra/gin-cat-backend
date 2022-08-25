package controllers

import (
	"net/http"

	"example.com/golang-restfulapi/models"
	"github.com/gin-gonic/gin"
)

// GET
// func (db *DBController) GetCats(c *gin.Context) {
// 	_type := c.Query("type")
// 	_where := map[string]interface{}{}

// 	if _type != "" {
// 		_where["type"] = _type
// 	}

// 	var cats []models.Cats
// 	db.Database.Where(_where).Find(&cats)

// 	c.JSON(http.StatusOK, gin.H{"results": &cats})
// }

// GET BY ID
// func (db *DBController) GetCollectionById(c *gin.Context) {
// 	id := c.Param("id")
// 	var collections models.Collections

// 	db.Database.First(&collections, id)
// 	db.Database.Model(&collections).Association("Groups").Find(&collections.Groups)

// 	c.JSON(http.StatusOK, gin.H{"results": &collections})
// }

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
// func (db *DBController) UpdateCollection(c *gin.Context) {

// 	var collection models.Collections
// 	err := c.ShouldBind(&collection)

// 	result := db.Database.Updates(collection)

// 	if result.Error != nil || err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"results": &collection})
// 	}
// }

// // DELETE
// func (db *DBController) DeleteCollection(c *gin.Context) {
// 	id := c.Param("id")
// 	var collections models.Collections
// 	db.Database.Delete(&collections, id)

// 	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK})
// }
