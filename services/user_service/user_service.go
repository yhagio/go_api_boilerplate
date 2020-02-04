package user_service

import (
	"errors"
	"go_api_boilerplate/domain/user"
	"go_api_boilerplate/repositories/user_repo"
)

type UserService interface {
	GetById(id int) (*user.User, error)
}

type userService struct {
	Repo user_repo.UserRepo
}

// NewUserService will instantiate User Service
func NewUserService(repo user_repo.UserRepo) UserService {
	return &userService{
		Repo: repo,
	}
}

func (us *userService) GetById(id int) (*user.User, error) {
	if id == 0 {
		return nil, errors.New("id param is required")
	}
	user, err := us.Repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
