package repository

import (
	"errors"
	"hendralijaya/austin-hendra-restapi/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(u domain.User) domain.User
	Update(u domain.User) domain.User
	Delete(u domain.User)
	FindByEmail(email string) (domain.User, error)
	VerifyCredential(email, password string) (domain.User, error)
	IsDuplicateEmail(email string) (bool, error)
}

type UserConnection struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &UserConnection{connection: connection}
}

func (c *UserConnection) Insert(u domain.User) domain.User {
	c.connection.Create(&u)
	c.connection.Find(&u)
	return u
}

func (c *UserConnection) Update(u domain.User) domain.User {
	c.connection.Save(&u)
	c.connection.Find(&u)
	return u
}

func (c *UserConnection) Delete(u domain.User) {
	c.connection.Delete(&u)
}

func (c *UserConnection) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	c.connection.Find(&user, "email = ?", email)
	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (c *UserConnection) VerifyCredential(email, password string) (domain.User, error) {
	var user domain.User
	c.connection.Find(&user, "email = ? AND password = ?", email, password)
	if user.Id == 0 {
		return user, errors.New("wrong credential")
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

