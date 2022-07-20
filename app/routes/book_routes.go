package routes

import (
	"hendralijaya/austin-hendra-restapi/controller"
	"hendralijaya/austin-hendra-restapi/middleware"
	"hendralijaya/austin-hendra-restapi/repository"
	"hendralijaya/austin-hendra-restapi/service"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

func NewBookRoutes(db *gorm.DB, route *gin.Engine) {
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	bookRoute := route.Group("/api/v1")
	bookRoute.Use(middleware.ErrorHandler)
	bookRoute.Use(cors.Default())
	bookRoute.GET("/book", bookController.All)
	bookRoute.GET("/book/:bookId", bookController.FindById)
	bookRoute.POST("/book", bookController.Insert)
	bookRoute.PUT("/book/:bookId", bookController.Update)
	bookRoute.DELETE("/book/:bookId", bookController.Delete)
}