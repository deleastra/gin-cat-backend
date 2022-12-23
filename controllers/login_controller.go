package controllers

import (
	"net/http"

	"cat-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginController struct {
	DB *gorm.DB
}

func NewLoginController(db *gorm.DB) *LoginController {
	return &LoginController{DB: db}
}

// @Summary Logs a user in
// @Description Logs a user in and returns a JWT token
// @Tags login
// @Accept json
// @Produce json
// @Param login body models.Login true "Credentials for logging in"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /login [post]
func (ctrl LoginController) Login(ctx *gin.Context) {
	// Parse the request and bind the login struct.
	var login models.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Retrieve the user from the database.
	var user models.User
	if err := ctrl.DB.Where("email = ?", login.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Message: "invalid email or password",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Check if the password is correct.
	if !user.CheckPassword(login.Password) {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Message: "invalid email or password",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Generate a JWT token.
	token, err := user.GenerateJWT()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Return the token to the client.
	ctx.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    token,
		Code:    http.StatusOK,
	})
}

// @Summary Logs a user out
// @Description Logs a user out by invalidating their JWT token
// @Tags login
// @Produce json
// @Success 200 {object} models.Response
// @Router /logout [post]
func (ctrl LoginController) Logout(ctx *gin.Context) {
	// Invalidate the JWT token by setting the expiration time to a time in the past
	ctx.SetCookie("jwt_token", "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, models.Response{
		Message: "Successfully logged out",
		Code:    http.StatusOK,
	})
}
