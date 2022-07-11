package service

import (
	"fmt"
	"hendralijaya/austin-hendra-restapi/helper"
	"hendralijaya/austin-hendra-restapi/model/domain"
	"hendralijaya/austin-hendra-restapi/model/web"
	"hendralijaya/austin-hendra-restapi/repository"
	"strconv"

	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(b web.BookCreateRequest) domain.Book
	Update(b web.BookUpdateRequest) domain.Book
	Delete(b domain.Book)
	FindById(bookId uint64) domain.Book
	All() []domain.Book
	IsAllowedToEdit(writerId uint64, bookId uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService (bookRepository repository.BookRepository) BookService {
	return &bookService{bookRepository: bookRepository}
}

func (service *bookService) Insert(b web.BookCreateRequest) domain.Book {
	book := domain.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	helper.PanicIfError(err)
	return service.bookRepository.Insert(book)
}

func (service *bookService) Update(b web.BookUpdateRequest) domain.Book {
	book := domain.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	helper.PanicIfError(err)
	return service.bookRepository.Update(book)
}

func (service *bookService) Delete(b domain.Book) {
	service.bookRepository.Delete(b)
}

func (service *bookService) FindById(bookId uint64) domain.Book {
	return service.bookRepository.FindById(bookId)
}

func (service *bookService) All() []domain.Book {
	return service.bookRepository.All()
}

func (service *bookService) IsAllowedToEdit(writerId uint64, bookId uint64) bool {
	book := service.bookRepository.FindById(bookId)
	id := fmt.Sprintf("%v", book.WriterId)
	return id == strconv.FormatUint(writerId, 10)
}