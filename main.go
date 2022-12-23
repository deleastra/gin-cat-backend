package main

import (
	"cat-backend/models"
	"cat-backend/routers"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	docs "cat-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Cat API
// @version 1.0
// @description A simple API for managing cats.

// @contact.name API Support
// @contact.email support@example.com

// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Connect to the database.
	db, err := gorm.Open(sqlite.Open("cat.db"))
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	db.AutoMigrate(&models.Cats{}, &models.Users{})

	// Set up the router.
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	// Set up the routes for the cat controller.
	routers.SetCatRoutes(db, r)

	// Set up Swagger documentation.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Run the server.
	r.Run()
}
