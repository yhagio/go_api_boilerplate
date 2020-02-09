package userservice

import (
	"errors"
	"go_api_boilerplate/domain/user"
	"go_api_boilerplate/repositories/userrepo"

	"golang.org/x/crypto/bcrypt"
)

// UserService interface
type UserService interface {
	GetByID(id uint) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	Create(*user.User) error
	Update(*user.User) error
	HashPassword(rawPassword string) (string, error)
	ComparePassword(rawPassword string, passwordFromDB string) error
}

type userService struct {
	Repo   userrepo.UserRepo
	pepper string
}

// NewUserService will instantiate User Service
func NewUserService(repo userrepo.UserRepo, pepper string) UserService {
	return &userService{
		Repo:   repo,
		pepper: pepper,
	}
}

func (us *userService) GetByID(id uint) (*user.User, error) {
	if id == 0 {
		return nil, errors.New("id param is required")
	}
	user, err := us.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) GetByEmail(email string) (*user.User, error) {
	if email == "" {
		return nil, errors.New("email(string) is required")
	}
	user, err := us.Repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) Create(user *user.User) error {
	hashedPass, err := us.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPass
	return us.Repo.Create(user)
}

func (us *userService) Update(user *user.User) error {
	return us.Repo.Update(user)
}

func (us *userService) HashPassword(rawPassword string) (string, error) {
	passAndPepper := rawPassword + us.pepper
	hashed, err := bcrypt.GenerateFromPassword([]byte(passAndPepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), err
}

func (us *userService) ComparePassword(rawPassword string, passwordFromDB string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(passwordFromDB),
		[]byte(rawPassword+us.pepper),
	)
}
