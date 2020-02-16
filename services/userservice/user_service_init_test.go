package userservice

import (
	"github.com/stretchr/testify/mock"
	"github.com/yhagio/go_api_boilerplate/domain/passwordreset"
	"github.com/yhagio/go_api_boilerplate/domain/user"
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

type pwdRepoMock struct {
	mock.Mock
}

func (repo *pwdRepoMock) Create(pwd *passwordreset.PasswordReset) error {
	args := repo.Called(pwd)
	return args.Error(0)
}

func (repo *pwdRepoMock) GetOneByToken(token string) (*passwordreset.PasswordReset, error) {
	args := repo.Called(token)
	return args.Get(0).(*passwordreset.PasswordReset), args.Error(1)
}

func (repo *pwdRepoMock) Delete(id uint) error {
	args := repo.Called(id)
	return args.Error(0)
}

type rdm struct{}

func (rd *rdm) GenerateToken() (string, error) {
	return "token", nil
}
func (rd *rdm) NumberOfBytes(base64String string) (int, error) {
	return 44, nil
}

type hmacMock struct{}

func (h *hmacMock) Hash(input string) string {
	return input + "hashed"
}
