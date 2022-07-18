package controller

import (
	"errors"
	"fmt"
	"hendralijaya/austin-hendra-restapi/helper"
	"hendralijaya/austin-hendra-restapi/model/web"
	"hendralijaya/austin-hendra-restapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WriterController interface {
	All(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type writerController struct {
	writerService service.WriterService
}

func NewWriterController(writerService service.WriterService) WriterController {
	return &writerController{writerService: writerService}
}

func (c *writerController) All(ctx *gin.Context) {
	writers := c.writerService.All()
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   writers,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *writerController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("writerId"), 10, 64)
	fmt.Println(id)
	if err != nil {
		helper.NotFoundError(ctx, errors.New("writer not found"))
	}
	writer, err := c.writerService.FindById(id)
	ok := helper.NotFoundError(ctx, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   writer,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *writerController) Insert(ctx *gin.Context) {
	var w web.WriterCreateRequest
	err := ctx.BindJSON(&w)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	writer, err := c.writerService.Insert(w)
	ok = helper.ValidationError(ctx, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success",
		Errors: nil,
		Data:   writer,
	}
	ctx.JSON(http.StatusCreated, webResponse)
}

func (c *writerController) Update(ctx *gin.Context) {
	var w web.WriterUpdateRequest
	err := ctx.BindJSON(&w)
	ok := helper.ValidationError(ctx, err)
	if ok {
		return
	}
	writer, err := c.writerService.Update(w)
	if err != nil {
		ctx.Error(err).SetMeta("NOT_FOUND")
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   writer,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *writerController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("writerId"), 10, 64)
	if err != nil {
		helper.NotFoundError(ctx, errors.New("writer not found"))
	}
	err = c.writerService.Delete(id)
	ok := helper.NotFoundError(ctx, err)
	if ok {
		return
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}
