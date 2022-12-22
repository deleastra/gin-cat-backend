package controllers

import (
	"net/http"
	"path/filepath"

	"cat-backend/models"
	"crypto/sha256"
	"encoding/hex"
	"io"

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

func (c *CatsController) CreateCat(ctx *gin.Context) {
	// Parse the request and bind the cat struct.
	var cat models.Cats
	cat.Name = ctx.Request.FormValue("name")

	// Read the image file from the request.
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Open the image file.
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a hash value for the image file.
	h := sha256.New()
	if _, err := io.Copy(h, src); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash := hex.EncodeToString(h.Sum(nil))

	// Rename the image file to the hash value.
	cat.Image = hash + filepath.Ext(file.Filename)

	// Upload the image file to the server.
	if err := ctx.SaveUploadedFile(file, "images/"+cat.Image); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the cat to the database.
	if err := c.DB.Create(&cat).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"cat": cat})
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
