package service

import (
	"hendralijaya/austin-hendra-restapi/exception"
	"hendralijaya/austin-hendra-restapi/model/domain"
	"hendralijaya/austin-hendra-restapi/model/web"
	"hendralijaya/austin-hendra-restapi/repository"

	"github.com/go-playground/validator/v10"
	"github.com/mashingan/smapping"
)

type WriterService interface {
	Insert(b web.WriterCreateRequest) (domain.Writer, error)
	Update(b web.WriterUpdateRequest) (domain.Writer, error)
	Delete(writerId uint64) error
	FindById(writerId uint64) (domain.Writer, error)
	All() []domain.Writer
	// IsAllowedToEdit(writerId uint64, bookId uint64) bool
}

type writerService struct {
	writerRepository repository.WriterRepository
	Validate         *validator.Validate
}

func NewWriterService(writerRepository repository.WriterRepository) WriterService {
	return &writerService{writerRepository: writerRepository}
}

func (service *writerService) Insert(request web.WriterCreateRequest) (domain.Writer, error) {
	writer := domain.Writer{}
	err := smapping.FillStruct(&writer, smapping.MapFields(&request))
	if err != nil {
		return writer, err
	}
	return service.writerRepository.Insert(writer), nil
}

func (service *writerService) Update(request web.WriterUpdateRequest) (domain.Writer, error) {
	writerRequest := domain.Writer{}
	err := smapping.FillStruct(&writerRequest, smapping.MapFields(&request))
	if err != nil {
		return writerRequest, err
	}
	_, err = service.writerRepository.FindById(request.Id)
	if err != nil {
		return writerRequest, exception.NewNotFoundError(err.Error())
	}
	return service.writerRepository.Update(writerRequest), nil
}

func (service *writerService) Delete(writerId uint64) error {
	writer, err := service.writerRepository.FindById(writerId)
	if err != nil {
		return exception.NewNotFoundError(err.Error())
	}
	service.writerRepository.Delete(writer)
	return nil
}

func (service *writerService) FindById(writerId uint64) (domain.Writer, error) {
	writer, err := service.writerRepository.FindById(writerId)
	if err != nil {
		return writer, err
	}
	return writer, nil
}

func (service *writerService) All() []domain.Writer {
	return service.writerRepository.All()
}

// func (service *writerService) IsAllowedToEdit(writerId uint64, bookId uint64) bool {
// 	writer, err := service.writerRepository.FindById(writerId)
// 	if err != nil {
// 		panic(exception.NewNotFoundError(err.Error()))
// 	}
// 	id := fmt.Sprintf("%v", writer.BookId)
// 	return id == strconv.FormatUint(bookId, 10)
// }
