package user_service

import (
	"errors"
	"go_api_boilerplate/domain/user"
	"go_api_boilerplate/repositories/user_repo"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetByID(id uint) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	Create(*user.User) error
	ComparePassword(rawPassword string, passwordFromDB string) error
}

type userService struct {
	Repo   user_repo.UserRepo
	pepper string
}

// NewUserService will instantiate User Service
func NewUserService(repo user_repo.UserRepo, pepper string) UserService {
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
	hashedPass, err := us.hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPass
	return us.Repo.Create(user)
}

func (us *userService) hashPassword(rawPassword string) (string, error) {
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
