package test

import (
	"encoding/json"
	"fmt"
	"hendralijaya/austin-hendra-restapi/app/routes"
	"hendralijaya/austin-hendra-restapi/helper"
	"hendralijaya/austin-hendra-restapi/middleware"
	"hendralijaya/austin-hendra-restapi/model/domain"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	db *gorm.DB = setupTestDB()
)

func setupTestDB() *gorm.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAMETEST")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	db.AutoMigrate(&domain.Book{}, &domain.Writer{})
	return db
}

func CloseTestDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	helper.PanicIfError(err)
	dbSQL.Close()
}

func TruncateTable(db *gorm.DB, tableName string) {
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", tableName))
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

func TestCreateBookSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	TruncateTable(db, "books")

	router := SetupRouter()

	requestBody := strings.NewReader(`{
		"title":"Belajar Golang Bersama Programmer Zaman Now",
		"publisher":"Eko Kurniawan Khannedy",
		"book_type":"E-Book",
		"year_released" : "2019",
		"synopsis":"Belajar Golang Bersama Programmer Zaman Now",
		"genre":"Programming",
		"stock":1,
		"writer_id":1
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/book", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusCreated)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusCreated, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, "Belajar Golang Bersama Programmer Zaman Now", responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, "Eko Kurniawan Khannedy", responseBody["data"].(map[string]interface{})["publisher"])
	fmt.Println()
}

func TestCreateFailed(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)

	router := SetupRouter()

	requestBody := strings.NewReader(`{
		"title":"Belajar Golang Bersama Programmer Zaman Now",
		"publisher":"Eko Kurniawan Khannedy",
		"book_type":"E-Book",
		"year_released" : "2019",
		"synopsis":"Belajar Golang Bersama Programmer Zaman Now",
		"genre":"",
		"stock":1,
		"writer_id":1
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/book", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusBadRequest)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
	assert.Equal(t, "Genre", responseBody["errors"].([]interface{})[0].(map[string]interface{})["key"])
	assert.Equal(t, "Error : failed on the required tag", responseBody["errors"].([]interface{})[0].(map[string]interface{})["message"])
}

func TestFindByIdSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	router := SetupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/book/1", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusOK)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, "Belajar Golang Bersama Programmer Zaman Now", responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, "Eko Kurniawan Khannedy", responseBody["data"].(map[string]interface{})["publisher"])
}

func TestFindByIdFailed(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	router := SetupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/book/100", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusNotFound)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
	assert.Equal(t, "book not found", responseBody["errors"].(map[string]interface{})["error"])
}

func TestUpdateSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"id":1,
		"title":"Belajar Golang Bersama Programmer Zaman Now",
		"publisher":"Eko Kurniawan Khannedy",
		"book_type":"E-Book",
		"year_released" : "2019",
		"synopsis":"Belajar Golang Bersama Programmer Zaman Now",
		"genre":"Programming",
		"stock":1,
		"writer_id":1
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/book/1", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusOK)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, "Belajar Golang Bersama Programmer Zaman Now", responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, "Eko Kurniawan Khannedy", responseBody["data"].(map[string]interface{})["publisher"])
}

func TestUpdateFailed(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	router := SetupRouter()
	requestBody := strings.NewReader(`{
		"id":100,
		"title":"Belajar Golang Bersama Programmer Zaman Now",
		"publisher":"Eko Kurniawan Khannedy",
		"book_type":"E-Book",
		"year_released" : "2019",
		"synopsis":"Belajar Golang Bersama Programmer Zaman Now",
		"genre":"Programming",
		"stock":1,
		"writer_id":1
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8000/api/v1/book/100", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusNotFound)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
	assert.Equal(t, "book not found", responseBody["errors"].(map[string]interface{})["error"])
}

func TestDeleteSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	router := SetupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/book/1", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t,int(responseBody["code"].(float64)), http.StatusOK)
	assert.Equal(t, responseBody["status"], "Success")
}

func TestDeleteFailed(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	router := SetupRouter()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8000/api/v1/book/100", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusNotFound)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
	assert.Equal(t, "book not found", responseBody["errors"].(map[string]interface{})["error"])
}

func TestFindAllSuccess(t *testing.T) {
	db := setupTestDB()
	defer CloseTestDB(db)
	router := SetupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000/api/v1/book", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, response.StatusCode, http.StatusOK)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "Success", responseBody["status"])
	assert.Equal(t, "Belajar Golang Bersama Programmer Zaman Now", responseBody["data"].([]interface{})[0].(map[string]interface{})["title"])
	assert.Equal(t, "Eko Kurniawan Khannedy", responseBody["data"].([]interface{})[0].(map[string]interface{})["publisher"])
	fmt.Println(responseBody)
}


