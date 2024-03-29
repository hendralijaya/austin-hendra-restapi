package service

import (
	"hendralijaya/austin-hendra-restapi/model/domain"
	"hendralijaya/austin-hendra-restapi/model/web"
	"hendralijaya/austin-hendra-restapi/repository"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	VerifyCredential(b web.UserLoginRequest) (interface{}, error)
	Create(b web.UserRegisterRequest) (domain.User, error)
	FindById(id uint64) (domain.User, error)
	Update(b domain.User) (domain.User, error)
	// Logout(u web.UserLogoutRequest) (domain.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (s *authService) VerifyCredential(u web.UserLoginRequest) (interface{}, error) {
	user, err := s.userRepository.VerifyCredential(u.Username, u.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *authService) Create(request web.UserRegisterRequest) (domain.User, error) {
	user := domain.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&request))

	if err != nil {
		return user, err
	}

	_, err = s.userRepository.FindByEmail(request.Email)
	if err != nil {
		return user, err
	}

	return s.userRepository.Create(user), nil
}

func (s *authService) Update(b domain.User) (domain.User, error) {
	_,err := s.userRepository.FindById(b.Id)
	if err != nil {
		return b, err
	}
	return s.userRepository.Update(b), nil
}

func (s *authService) FindById(id uint64) (domain.User, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
