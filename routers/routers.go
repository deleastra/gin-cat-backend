package routers

import (
	"example.com/golang-restfulapi/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetCatRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ctrls := controllers.DBController{Database: db}

	// router.GET("collections", ctrls.GetCollection)  // GET
	// router.GET("collections/:id", ctrls.GetCollectionById)  // GET BY ID
	router.POST("cat", ctrls.CreateCat) // POST
	// router.PATCH("collections", ctrls.UpdateCollection)  // PATCH
	// router.DELETE("collections/:id", ctrls.DeleteCollection)  // DELETE
}
