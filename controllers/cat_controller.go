package controllers

import (
	"crypto/rand"
	"crypto/sha512"
	"image"
	"image/color"
	"net/http"
	"path/filepath"

	"cat-backend/models"
	"encoding/hex"

	"github.com/disintegration/imaging"

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

	// Decode the image into memory.
	srcImg, err := imaging.Decode(src)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Resize the image to a specific width and height.
	dstImg := imaging.Resize(srcImg, 800, 600, imaging.Lanczos)

	// Create a new image file in memory to store the resized image.
	dst := imaging.New(800, 600, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, dstImg, image.Pt(0, 0))

	// Generate a random slice of bytes.
	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Generate the SHA-512 hash of the random bytes.
	hash := sha512.Sum512(randomBytes)

	// Encode the hash as a hexadecimal string.
	hashString := hex.EncodeToString(hash[:])

	// Rename the image file to the hash value.
	cat.Image = hashString + filepath.Ext(file.Filename)

	// Upload the image file to the server.
	if imaging.Save(dst, "images/"+cat.Image); err != nil {
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
