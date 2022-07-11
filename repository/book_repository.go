package repository

import (
	"hendralijaya/austin-hendra-restapi/model/domain"

	"gorm.io/gorm"
)

type BookRepository interface {
	All() []domain.Book
	Insert(b domain.Book) domain.Book
	Update(b domain.Book) domain.Book
	Delete(b domain.Book)
	FindById(bookId uint64) domain.Book
}

type BookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(connection *gorm.DB) BookRepository {
	return &BookConnection{connection: connection}
}

func(db *BookConnection) All() []domain.Book {
	var books []domain.Book
	db.connection.Preload("Writer").Find(&books)
	return books
}

func (db *BookConnection) Insert(book domain.Book) domain.Book {
	db.connection.Save(&book)
	db.connection.Preload("Writer").Find(&book)
	return book
}

func (db *BookConnection) Update(book domain.Book) domain.Book {
	db.connection.Save(&book)
	db.connection.Preload("Writer").Find(&book)
	return book
}

func (db *BookConnection) Delete(book domain.Book) {
	db.connection.Delete(&book)
}

func (db *BookConnection) FindById(id uint64) domain.Book {
	var book domain.Book
	db.connection.Preload("Writer").Find(&book, id)
	return book
}