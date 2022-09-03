package main

import (
	"fmt"
	"net/http"

	"example.com/golang-restfulapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	routers "example.com/golang-restfulapi/routers"
)

func main() {
	db, err := gorm.Open(sqlite.Open("cat.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var cat models.Cats
	var user models.Users
	// Migrate the schema
	db.AutoMigrate(cat)
	db.AutoMigrate(user)

	router := gin.New()
	api := router.Group("/cat")
	routers.SetCatRoutes(api, db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO GOLANG RESTFUL API.",
		})
	})
	fmt.Println("Server Running on Port: ", 5000)
	http.ListenAndServe(":5000", router)
}
