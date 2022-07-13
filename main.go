package main

import (
	"hendralijaya/austin-hendra-restapi/app/config"
	"hendralijaya/austin-hendra-restapi/app/routes"
	"hendralijaya/austin-hendra-restapi/helper"
	"hendralijaya/austin-hendra-restapi/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/gorm"
)

var(
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	err := godotenv.Load()
	helper.PanicIfError(err)
	router := SetupRouter()
	log.Fatal(router.Run(":" + os.Getenv("GO_PORT")))
}

func SetupRouter() *gin.Engine {
	err := godotenv.Load()
	helper.PanicIfError(err)
	/**
	@description Setup Database Connection
	*/
	
	/**
	@description Init Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	if os.Getenv("GO_ENV") != "production" && os.Getenv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if os.Getenv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	@description Setup Middleware
	*/
	
	/**
	@description Init All Route
	*/
	routes.NewBookRoutes(db, router)
	router.Use(middleware.ErrorHandler)
	router.Use(cors.Default())

	return router
}