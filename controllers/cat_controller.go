package controllers

import (
	"net/http"

	"cat-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CatsController represents a controller for handling HTTP requests related to cats.
type CatsController struct {
	DB *gorm.DB
}

func NewCatsController(db *gorm.DB) *CatsController {
	return &CatsController{DB: db}
}

// GetCats handles a GET request to retrieve a list of all cats.
func (ctrl CatsController) GetCats(c *gin.Context) {
	var cats []models.Cats
	if err := ctrl.DB.Find(&cats).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, cats)
}

// GetCatByID handles a GET request to retrieve a single cat by ID.
func (ctrl CatsController) GetCatByID(c *gin.Context) {
	var cat models.Cats
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&cat).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, cat)
}

// CreateCat handles a POST request to create a new cat.
func (ctrl CatsController) CreateCat(c *gin.Context) {
	var cat models.Cats
	if err := c.BindJSON(&cat); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := ctrl.DB.Create(&cat).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, cat)
}

// UpdateCat handles a PUT request to update an existing cat.
func (ctrl CatsController) UpdateCat(c *gin.Context) {
	var cat models.Cats
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&cat).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err := c.BindJSON(&cat); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := ctrl.DB.Save(&cat).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, cat)
}

// DeleteCat handles a DELETE request to delete an existing cat.
func (ctrl CatsController) DeleteCat(c *gin.Context) {
	var cat models.Cats
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&cat).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err := ctrl.DB.Delete(&cat).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
