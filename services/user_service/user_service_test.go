package user_service

import (
	"errors"
	"go_api_boilerplate/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	pepper    = "pepper"
	testID10  = uint(10)
	testID100 = uint(100)
)

type repoMock struct {
	mock.Mock
}

func (repo *repoMock) GetByID(id uint) (*user.User, error) {
	args := repo.Called(id)
	return args.Get(0).(*user.User), args.Error(1)
}

func (repo *repoMock) GetByEmail(email string) (*user.User, error) {
	args := repo.Called(email)
	return args.Get(0).(*user.User), args.Error(1)
}

func (repo *repoMock) Create(user *user.User) error {
	args := repo.Called(user)
	return args.Error(1)
}

func (repo *repoMock) Update(user *user.User) error {
	args := repo.Called(user)
	return args.Error(1)
}

func TestGetByID(t *testing.T) {
	t.Run("Get a user", func(t *testing.T) {
		expected := &user.User{
			FirstName: "Test",
			LastName:  "User",
		}

		userRepo := new(repoMock)
		userRepo.On("GetByID", testID100).Return(expected, nil)
		u := NewUserService(userRepo, pepper)
		result, _ := u.GetByID(testID100)

		assert.EqualValues(t, expected, result)
	})

	t.Run("Get error if id is 0", func(t *testing.T) {
		expected := errors.New("id param is required")

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)
		result, err := u.GetByID(0)

		assert.Nil(t, result)
		assert.EqualValues(t, expected, err)
	})

	t.Run("Get error if it has error", func(t *testing.T) {
		expected := errors.New("Nop")

		userRepo := new(repoMock)
		userRepo.On("GetByID", testID10).Return(&user.User{}, expected)
		u := NewUserService(userRepo, pepper)
		result, err := u.GetByID(testID10)

		assert.Nil(t, result)
		assert.EqualValues(t, expected, err)
	})
}
