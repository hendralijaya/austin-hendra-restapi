package repository

import (
	"errors"
	"hendralijaya/austin-hendra-restapi/model/domain"

	"gorm.io/gorm"
)

type WriterRepository interface {
	All() []domain.Writer
	Insert(b domain.Writer) domain.Writer
	Update(b domain.Writer) domain.Writer
	Delete(b domain.Writer)
	FindById(writerId uint64) (domain.Writer, error)
}

type WriterConnection struct {
	connection *gorm.DB
}

func NewWriterRepository(connection *gorm.DB) WriterRepository {
	return &WriterConnection{connection: connection}
}

func (db *WriterConnection) All() []domain.Writer {
	var writers []domain.Writer
	db.connection.Preload("Writer").Find(&writers)
	return writers
}

func (db *WriterConnection) Insert(writer domain.Writer) domain.Writer {
	db.connection.Save(&writer)
	db.connection.Preload("Writer").Find(&writer)
	return writer
}

func (db *WriterConnection) Update(writer domain.Writer) domain.Writer {
	db.connection.Save(&writer)
	db.connection.Preload("Writer").Find(&writer)
	return writer
}

func (db *WriterConnection) Delete(writer domain.Writer) {
	db.connection.Delete(&writer)
}

func (db *WriterConnection) FindById(id uint64) (domain.Writer, error) {
	var writer domain.Writer
	db.connection.Preload("Writer").Find(&writer, id)
	if writer.Id == 0 {
		return writer, errors.New("writer not found")
	}
	return writer, nil
}
