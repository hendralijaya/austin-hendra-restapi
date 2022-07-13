package service

import (
	"fmt"
	"hendralijaya/austin-hendra-restapi/exception"
	"hendralijaya/austin-hendra-restapi/model/domain"
	"hendralijaya/austin-hendra-restapi/model/web"
	"hendralijaya/austin-hendra-restapi/repository"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(b web.BookCreateRequest) (domain.Book,error)
	Update(b web.BookUpdateRequest) (domain.Book, error)
	Delete(bookId uint64) error
	FindById(bookId uint64) (domain.Book, error)
	All() []domain.Book
	IsAllowedToEdit(writerId uint64, bookId uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
	Validate *validator.Validate
}

func NewBookService (bookRepository repository.BookRepository) BookService {
	return &bookService{bookRepository: bookRepository}
}

func (service *bookService) Insert(request web.BookCreateRequest) (domain.Book, error) {
	book := domain.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&request))
	if(err != nil) {
		return book, err
	}
	return service.bookRepository.Insert(book), nil
}

func (service *bookService) Update(request web.BookUpdateRequest) (domain.Book, error) {
	bookRequest := domain.Book{}
	err := smapping.FillStruct(&bookRequest, smapping.MapFields(&request))
	if(err != nil) {
		return bookRequest, err
	}
	_ ,err = service.bookRepository.FindById(request.Id)
	if err != nil {
		return bookRequest, exception.NewNotFoundError(err.Error())
	}
	return service.bookRepository.Update(bookRequest), nil
}

func (service *bookService) Delete(bookId uint64) error {
	book , err := service.bookRepository.FindById(bookId)
	if err != nil {
		return exception.NewNotFoundError(err.Error())
	}
	service.bookRepository.Delete(book)
	return nil
}

func (service *bookService) FindById(bookId uint64) (domain.Book, error) {
	book, err := service.bookRepository.FindById(bookId)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (service *bookService) All() []domain.Book {
	return service.bookRepository.All()
}

func (service *bookService) IsAllowedToEdit(writerId uint64, bookId uint64) bool {
	book, err := service.bookRepository.FindById(bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	id := fmt.Sprintf("%v", book.WriterId)
	return id == strconv.FormatUint(writerId, 10)
}