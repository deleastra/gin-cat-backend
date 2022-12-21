package main

import (
	"log"
	"net/http"

	"cat-backend/routers"

	"github.com/gin-gonic/gin"
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
	db.AutoMigrate()
	// Set up router.
	router := gin.Default()
	r := routers.SetCatRoutes(db, router)

	// Start server.
	// Enable debugging.
	log.Println("server listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
