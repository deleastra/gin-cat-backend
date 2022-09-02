package routers

import (
	"example.com/golang-restfulapi/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetCatRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ctrls := controllers.DBController{Database: db}

	router.GET("cat", ctrls.GetCats)  // GET
	router.GET("cat/:id", ctrls.GetCatByID)  // GET BY ID
	router.POST("cat", ctrls.CreateCat) // POST
	router.PATCH("cat", ctrls.UpdateCat)  // PATCH
	router.DELETE("cat/:id", ctrls.DeleteCat)  // DELETE
}
