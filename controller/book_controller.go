package controller

import (
	"hendralijaya/austin-hendra-restapi/helper"
	"hendralijaya/austin-hendra-restapi/model/domain"
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
	helper.PanicIfError(err)
	book := c.bookService.FindById(id)
	if (book == domain.Book{}) {
		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "Not Found",
			Errors: nil,
			Data: nil,
		}
		ctx.JSON(http.StatusNotFound, webResponse)
	}
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
	helper.PanicIfError(err)
	book := c.bookService.Insert(b)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data: book,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *bookController) Update(ctx *gin.Context) {
	var b web.BookUpdateRequest
	err := ctx.BindJSON(&b)
	helper.PanicIfError(err)
	book := c.bookService.Update(b)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data: book,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *bookController) Delete(ctx *gin.Context) {
	var book domain.Book
	id, err := strconv.ParseUint(ctx.Param("bookId"), 10, 64)
	helper.PanicIfError(err)
	book.Id = id
	c.bookService.Delete(book)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success",
		Errors: nil,
		Data: nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}