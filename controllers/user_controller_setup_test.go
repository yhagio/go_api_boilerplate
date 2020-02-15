package controllers

import (
	"github.com/yhagio/go_api_boilerplate/domain/user"
	"github.com/yhagio/go_api_boilerplate/services/authservice"
)

type userSvc struct{}

var sampleUser = &user.User{
	Email:     "alice@cc.cc",
	FirstName: "",
	LastName:  "",
	Active:    false,
	Role:      "",
}

func (us *userSvc) GetByID(id uint) (*user.User, error) {
	return sampleUser, nil
}

func (us *userSvc) GetByEmail(email string) (*user.User, error) {
	return sampleUser, nil
}

func (us *userSvc) Create(user *user.User) error {
	return nil
}

func (us *userSvc) Update(user *user.User) error {
	return nil
}

func (us *userSvc) HashPassword(rawPassword string) (string, error) {
	return rawPassword, nil
}

func (us *userSvc) ComparePassword(rawPassword string, passwordFromDB string) error {
	return nil
}

type authSvc struct {
	jwtSecret string
}

func (auth *authSvc) IssueToken(u user.User) (string, error) {
	return "nice-token", nil
}

func (auth *authSvc) ParseToken(token string) (*authservice.Claims, error) {
	return nil, nil
}

type emailSvc struct{}

func (es *emailSvc) Welcome(toEmail string) error {
	return nil
}

func (es *emailSvc) ResetPassword(toEmail, token string) error {
	return nil
}
