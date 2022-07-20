package controller

import (
	"errors"
	"hendralijaya/austin-hendra-restapi/helper"
	"hendralijaya/austin-hendra-restapi/model/web"
	"hendralijaya/austin-hendra-restapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &bookController{bookService: bookService}
}

func (c *bookController) All(ctx *gin.Context) {
	books := c.bookService.All()
	webResponse := web.WebResponse{
		Code:  http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data:   books,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *bookController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("bookId"), 10, 64)
	if(err != nil) {
		helper.NotFoundError(ctx, errors.New("book not found"))
	}
	book, err := c.bookService.FindById(id)
	ok := helper.NotFoundError(ctx, err)
	if ok {return}
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data: book,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *bookController) Insert(ctx *gin.Context) {
	var b web.BookCreateRequest
	err := ctx.BindJSON(&b)
	ok := helper.ValidationError(ctx, err)
	if ok {return}
	book, err := c.bookService.Insert(b)
	ok = helper.ValidationError(ctx, err)
	if ok {return}
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "Success",
		Errors: nil,
		Data: book,
	}
	ctx.JSON(http.StatusCreated, webResponse)
}

func (c *bookController) Update(ctx *gin.Context) {
	var b web.BookUpdateRequest
	id, err := strconv.ParseUint(ctx.Param("bookId"), 10, 64)
	if(err != nil) {
		helper.NotFoundError(ctx, errors.New("book not found"))
	}
	b.Id = id
	err = ctx.BindJSON(&b)
	ok := helper.ValidationError(ctx, err)
	if ok {return}
	book, err := c.bookService.Update(b)
	if err != nil {
		ctx.Error(err).SetMeta("NOT_FOUND")
		return
	}
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data: book,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *bookController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("bookId"), 10, 64)
	if(err != nil) {
		helper.NotFoundError(ctx, errors.New("book not found"))
	}
	err = c.bookService.Delete(id)
	ok := helper.NotFoundError(ctx, err)
	if ok {return}
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data: nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}