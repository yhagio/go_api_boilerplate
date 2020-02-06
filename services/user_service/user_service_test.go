package user_service

import (
	"errors"
	"go_api_boilerplate/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repoMock struct {
	mock.Mock
}

func (repo *repoMock) GetById(id int) (*user.User, error) {
	args := repo.Called(id)

	return args.Get(0).(*user.User), args.Error(1)
}

func TestGetById(t *testing.T) {
	t.Run("Get a user", func(t *testing.T) {
		expected := &user.User{
			FirstName: "Test",
			LastName:  "User",
		}

		userRepo := new(repoMock)

		userRepo.On("GetById", 100).Return(expected, nil)

		u := NewUserService(userRepo)

		result, _ := u.GetById(100)

		assert.EqualValues(t, expected, result)
	})

	t.Run("Get error if id is 0", func(t *testing.T) {
		expected := errors.New("id param is required")

		userRepo := new(repoMock)

		u := NewUserService(userRepo)

		result, err := u.GetById(0)

		assert.Nil(t, result)
		assert.EqualValues(t, expected, err)
	})

	t.Run("Get error if it has error", func(t *testing.T) {
		expected := errors.New("Nop")

		userRepo := new(repoMock)
		userRepo.On("GetById", 10).Return(&user.User{}, expected)

		u := NewUserService(userRepo)

		result, err := u.GetById(10)

		assert.Nil(t, result)
		assert.EqualValues(t, expected, err)
	})
}
