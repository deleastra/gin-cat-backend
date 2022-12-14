package controllers

import (
	"crypto/rand"
	"crypto/sha512"
	"image"
	"image/color"
	"io/ioutil"
	"log"
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

// @Summary Retrieves a list of all cats
// @Description Retrieves a list of all cats stored in the database.
// @Tags cats
// @Produce json
// @Success 200 {array} models.Cats
// @Failure 500 {object} models.Response
// @Router /cats [get]
func (ctrl CatsController) GetCats(c *gin.Context) {
	var cats []models.Cats
	if err := ctrl.DB.Find(&cats).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, cats)
}

// @Summary Retrieves a single cat by ID
// @Description Retrieves a single cat by ID from the database.
// @Tags cats
// @Produce  json
// @Param id path int true "ID of the cat"
// @Success 200 {object} models.Cats
// @Failure 404 {object} models.Response
// @Router /cats/{id} [get]
func (ctrl CatsController) GetCatByID(c *gin.Context) {
	var cat models.Cats
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&cat).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, cat)
}

// @Summary Creates a new cat
// @Description Creates a new cat and stores it in the database.
// @Tags cats
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData string true "Name of the cat"
// @Param image formData file true "Image of the cat"
// @Success 200 {object} models.Cats
// @Failure 400 {object} models.Response
// @Router /cats [post]
func (c *CatsController) CreateCat(ctx *gin.Context) {
	// Parse the request and bind the cat struct.
	var cat models.Cats
	cat.Name = ctx.Request.FormValue("name")

	// Read the image file from the request.
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Open the image file.
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}
	defer src.Close()

	// Decode the image into memory.
	srcImg, err := imaging.Decode(src)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
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
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	// Generate the SHA-512 hash of the random bytes.
	hash := sha512.Sum512(randomBytes)

	// Encode the hash as a hexadecimal string.
	hashString := hex.EncodeToString(hash[:])

	// Rename the image file to the hash value.
	cat.Image = hashString + filepath.Ext(file.Filename)

	// Upload the image file to the server.
	if imaging.Save(dst, "images/"+cat.Image); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Save the cat to the database.
	if err := c.DB.Create(&cat).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"cat": cat})
}

// @Summary Updates a cat by ID
// @Description Updates a cat by ID and stores the changes in the database.
// @Tags cats
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the cat"
// @Param cat body models.Cats true "Updated cat information"
// @Success 200 {object} models.Cats
// @Failure 400 {object} models.Response
// @Router /cats/{id} [put]
func (c *CatsController) UpdateCat(ctx *gin.Context) {
	// Parse the request and bind the cat struct.
	var cat models.Cats
	if err := ctx.Bind(&cat); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Update the cat in the database.
	if err := c.DB.Model(&cat).Where("id = ?", ctx.Param("id")).Updates(&cat).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, cat)
}

// @Summary Deletes a cat by ID
// @Description Deletes a cat from the database by ID.
// @Tags cats
// @Produce  json
// @Param id path string true "ID of the cat"
// @Success 204
// @Failure 404 {object} models.Response
// @Router /cats/{id} [delete]
func (c *CatsController) DeleteCat(ctx *gin.Context) {
	var cat models.Cats
	// Delete the cat from the database.
	if err := c.DB.Where("id = ?", ctx.Param("id")).Delete(&models.Cats{}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		})
		return
	}
	if err := c.DB.Delete(&cat).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Summary Shows an image
// @Description Shows an image by ID.
// @Tags images
// @Produce image/jpeg
// @Param id path int true "ID of the image"
// @Success 200 {file} image/jpeg
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /images/{id} [get]
func (ctrl CatsController) ShowImage(c *gin.Context) {
	// Parse the cat ID from the request.
	catID := c.Param("id")

	// Look up the cat in the database using the ID.
	var cat models.Cats
	if err := ctrl.DB.Where("id = ?", catID).First(&cat).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Response{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		})
		return
	}

	log.Println(cat.Image)
	// Set the content type to the appropriate image type.
	c.Header("Content-Type", "image/jpeg")

	// Write the image data to the response body.
	 // Read the file into a byte slice
	 file, err := ioutil.ReadFile("images/" + cat.Image)
	 if err != nil {
		log.Fatalln(err)
		 c.JSON(http.StatusInternalServerError, models.Response{
			 Message: err.Error(),
			 Code:    http.StatusInternalServerError,
		 })
		 return
	 }
 
	 // Write the byte slice to the response body
	 c.Writer.Write(file)
}


