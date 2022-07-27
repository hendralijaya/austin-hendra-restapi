package routes

import (
	"hendralijaya/austin-hendra-restapi/controller"
	"hendralijaya/austin-hendra-restapi/helper"
	"hendralijaya/austin-hendra-restapi/middleware"
	"hendralijaya/austin-hendra-restapi/repository"
	"hendralijaya/austin-hendra-restapi/service"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

func NewAuthenticationRoutes(db *gorm.DB, route *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	authRoute := route.Group("/api/v1", helper.SetSession())
	authRoute.Use(middleware.ErrorHandler)
	authRoute.Use(cors.Default())
	authRoute.POST("/login/", authController.Login)
	authRoute.POST("/register/", authController.Register)
	authRoute.POST("/logout/", authController.Logout)
	authRoute.POST("/forgot_password/", authController.ForgotPassword)
}
