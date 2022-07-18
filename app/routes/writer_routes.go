package routes

import (
	"hendralijaya/austin-hendra-restapi/controller"
	"hendralijaya/austin-hendra-restapi/middleware"
	"hendralijaya/austin-hendra-restapi/repository"
	"hendralijaya/austin-hendra-restapi/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewWriterRoutes(db *gorm.DB, route *gin.Engine) {
	writerRepository := repository.NewWriterRepository(db)
	writerService := service.NewWriterService(writerRepository)
	writerController := controller.NewWriterController(writerService)

	writerRoute := route.Group("/api/v1")
	writerRoute.Use(middleware.ErrorHandler)
	writerRoute.GET("/writer", writerController.All)
	writerRoute.GET("/writer/:writerId", writerController.FindById)
	writerRoute.POST("/writer", writerController.Insert)
	writerRoute.PUT("/writer/:writerId", writerController.Update)
	writerRoute.DELETE("/writer/:writerId", writerController.Delete)
}
