package main

import (
	"log"
	"net/http"

	"cat-backend/models"
	"cat-backend/routers"

	docs "cat-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Set up database connection.
	db, err := gorm.Open(sqlite.Open("cat.db"))
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	db.AutoMigrate(&models.Cats{}, &models.Users{})
	// Set up router.
	router := gin.Default()
	r := routers.SetCatRoutes(db, router)
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/")
	{
		eg := v1.Group("/cat")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Start server.
	// Enable debugging.
	log.Println("server listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
