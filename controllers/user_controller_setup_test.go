package controllers

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/yhagio/go_api_boilerplate/domain/user"
	"github.com/yhagio/go_api_boilerplate/services/authservice"
)

type userSvc struct{}

var alice = &user.User{
	Model:     gorm.Model{ID: uint(1)},
	Email:     "alice@cc.cc",
	FirstName: "",
	LastName:  "",
	Active:    false,
	Role:      "",
}

var david = &user.User{
	Model:     gorm.Model{ID: uint(1)},
	Email:     "david@cc.cc",
	FirstName: "",
	LastName:  "",
	Active:    false,
	Role:      "",
}

func (us *userSvc) GetByID(id uint) (*user.User, error) {
	if id >= uint(100) {
		return nil, errors.New("Ugh")
	}
	if id < uint(1) {
		return nil, errors.New("Nop")
	}
	if id >= uint(10) {
		return nil, errors.New("Record not found")
	}
	return alice, nil
}

func (us *userSvc) GetByEmail(email string) (*user.User, error) {
	if email == "bob@cc.cc" {
		return nil, errors.New("Nop")
	}
	if email == david.Email {
		return david, nil
	}
	return alice, nil
}

func (us *userSvc) Create(user *user.User) error {
	if user.Email == "bob@cc.cc" {
		return errors.New("Nop")
	}
	return nil
}

func (us *userSvc) Update(user *user.User) error {
	if user.Email == "bob@cc.cc" {
		return errors.New("Nop")
	}
	return nil
}

func (us *userSvc) HashPassword(rawPassword string) (string, error) {
	return rawPassword, nil
}

func (us *userSvc) ComparePassword(rawPassword string, passwordFromDB string) error {
	if rawPassword != "123test" {
		return errors.New("Nop")
	}
	return nil
}

func (us *userSvc) InitiateResetPassowrd(email string) (string, error) {
	if email == "bob@cc.cc" {
		return "", errors.New("Nop")
	}
	return "token", nil
}

func (us *userSvc) CompleteUpdatePassword(token, newPassword string) (*user.User, error) {
	if newPassword == "xxx" {
		return nil, errors.New("Nop")
	}
	if newPassword == "david-pass" {
		return david, nil
	}
	return alice, nil
}

type authSvc struct {
	jwtSecret string
}

func (auth *authSvc) IssueToken(u user.User) (string, error) {
	if u.Email == david.Email {
		return "", errors.New("Nop")
	}
	return "nice-token", nil
}

func (auth *authSvc) ParseToken(token string) (*authservice.Claims, error) {
	return nil, nil
}

type emailSvc struct{}

func (es *emailSvc) Welcome(toEmail string) error {
	if toEmail == "chris@cc.cc" {
		return errors.New("Nop")
	}
	return nil
}

func (es *emailSvc) ResetPassword(toEmail, token string) error {
	if toEmail == "chris@cc.cc" {
		return errors.New("Nop")
	}
	return nil
}
