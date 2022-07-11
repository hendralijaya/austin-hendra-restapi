package middleware

import (
	"hendralijaya/austin-hendra-restapi/exception"
	"hendralijaya/austin-hendra-restapi/model/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(c *gin.Context, err error) {
	if validationErrors(c, err) {
		return
	}

	if notFoundError(c, err) {
		return
	}
	internalServerError(c, err)
}

func validationErrors(c *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		c.Header("Content-Type", "application/json")
		c.Status(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return true
	}else{
		return false
	}
}

func notFoundError(c *gin.Context, err interface{}) bool {
	exception, ok := err.(exception.NotFoundError)
	if ok {
		c.Header("Content-Type", "application/json")
		c.Status(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Errors: exception.Error,
		}
		c.JSON(http.StatusNotFound, webResponse)
		return true
	}else {
		return false
	}
}

func internalServerError(c *gin.Context, err interface{}) {
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	c.JSON(http.StatusInternalServerError, webResponse)
}