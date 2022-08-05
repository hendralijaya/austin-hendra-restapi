package repository

import (
	"errors"
	"hendralijaya/austin-hendra-restapi/helper"
	"hendralijaya/austin-hendra-restapi/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(u domain.User) domain.User
	// Update(u domain.User) domain.User
	// Delete(u domain.User)
	VerifyCredential(userName, password string) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	IsDuplicateEmail(email string) (bool, error)
}

type UserConnection struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &UserConnection{connection: connection}
}

func (c *UserConnection) Create(u domain.User) domain.User {
	u.Password = helper.HashAndSalt([]byte(u.Password))
	c.connection.Save(&u)
	return u
}

func (c *UserConnection) VerifyCredential(userName, password string) (domain.User, error) {
	var user domain.User
	c.connection.Find(&user, "username = ? AND password = ?", userName, password)
	if user.Id == 0 {
		return user, errors.New("wrong credential")
	}
	return user, nil
}

func (c *UserConnection) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	c.connection.Find(&user, "email = ?", email)
	if user.Id == 0 {
		return user, errors.New("email already exist")
	}
	return user, nil
}

func (c *UserConnection) IsDuplicateEmail(email string) (bool, error) {
	var user domain.User
	c.connection.Find(&user, "email = ?", email)
	if user.Id == 0 {
		return false, nil
	}
	return true, errors.New("email already exists")
}
