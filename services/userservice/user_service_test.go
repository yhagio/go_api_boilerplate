package userservice

import (
	"errors"
	"github.com/yhagio/go_api_boilerplate/domain/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	pepper    = "pepper"
	testID10  = uint(10)
	testID100 = uint(100)
	testEmail = "test@cc.cc"
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
	return args.Error(0)
}

func (repo *repoMock) Update(user *user.User) error {
	args := repo.Called(user)
	return args.Error(0)
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

func TestGetByEmail(t *testing.T) {
	t.Run("Get a user", func(t *testing.T) {
		expected := &user.User{
			FirstName: "Test",
			LastName:  "User",
		}

		userRepo := new(repoMock)
		userRepo.On("GetByEmail", testEmail).Return(expected, nil)
		u := NewUserService(userRepo, pepper)
		result, _ := u.GetByEmail(testEmail)

		assert.EqualValues(t, expected, result)
	})

	t.Run("Get error if email is empty", func(t *testing.T) {
		expected := errors.New("email(string) is required")

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)
		result, err := u.GetByEmail("")

		assert.Nil(t, result)
		assert.EqualValues(t, expected, err)
	})

	t.Run("Get error if it has error", func(t *testing.T) {
		expected := errors.New("Nop")

		userRepo := new(repoMock)
		userRepo.On("GetByEmail", testEmail).Return(&user.User{}, expected)
		u := NewUserService(userRepo, pepper)
		result, err := u.GetByEmail(testEmail)

		assert.Nil(t, result)
		assert.EqualValues(t, expected, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create a user", func(t *testing.T) {
		usr := &user.User{
			Email:    "alice@cc.cc",
			Password: "abc123",
		}

		userRepo := new(repoMock)
		userRepo.On("Create", usr).Return(nil)
		u := NewUserService(userRepo, pepper)
		result := u.Create(usr)

		assert.Nil(t, result)
	})

	t.Run("Create a user fails", func(t *testing.T) {
		err := errors.New(("oops"))
		usr := &user.User{
			Email: "alice@cc.cc",
		}

		userRepo := new(repoMock)
		userRepo.On("Create", usr).Return(err)
		u := NewUserService(userRepo, pepper)
		result := u.Create(usr)

		assert.EqualValues(t, result, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update a user", func(t *testing.T) {
		usr := &user.User{
			Email: "alice@cc.cc",
		}

		userRepo := new(repoMock)
		userRepo.On("Update", usr).Return(nil)
		u := NewUserService(userRepo, pepper)
		result := u.Update(usr)

		assert.Nil(t, result)
	})

	t.Run("Update a user fails", func(t *testing.T) {
		err := errors.New(("oops"))
		usr := &user.User{
			Email: "alice@cc.cc",
		}

		userRepo := new(repoMock)
		userRepo.On("Update", usr).Return(err)
		u := NewUserService(userRepo, pepper)
		result := u.Update(usr)

		assert.EqualValues(t, result, err)
	})
}

func TestComparePassword(t *testing.T) {
	t.Run("match password", func(t *testing.T) {
		testPass := "test123"

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)
		hashedPass, err := u.HashPassword(testPass)
		err = u.ComparePassword(testPass, hashedPass)
		assert.Nil(t, err)
	})

	t.Run("not match password", func(t *testing.T) {
		testPass := "test123"

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)
		hashedPass, err := u.HashPassword(testPass)
		err = u.ComparePassword("test1234", hashedPass)
		assert.NotNil(t, err)
	})
}
