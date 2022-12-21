package routers

import (
	"cat-backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetCatRoutes sets up the routes for the cat controller.
func SetCatRoutes(db *gorm.DB, r *gin.Engine) *gin.Engine {
	catCtrl := controllers.NewCatsController(db)

	r.GET("/cats", catCtrl.GetCats)
	r.POST("/cats", catCtrl.CreateCat)
	r.GET("/cats/:id", catCtrl.GetCatByID)
	r.PUT("/cats/:id", catCtrl.UpdateCat)
	r.DELETE("/cats/:id", catCtrl.DeleteCat)

	return r
}
