package userservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/yhagio/go_api_boilerplate/common/hmachash"
	rdms "github.com/yhagio/go_api_boilerplate/common/randomstring"
	pwd "github.com/yhagio/go_api_boilerplate/domain/passwordreset"
	"github.com/yhagio/go_api_boilerplate/domain/user"

	pwdRepo "github.com/yhagio/go_api_boilerplate/repositories/passwordreset"
	"github.com/yhagio/go_api_boilerplate/repositories/userrepo"

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
	InitiateResetPassowrd(email string) (string, error)
	CompleteUpdatePassword(token, newPassword string) (*user.User, error)
}

type userService struct {
	Repo    userrepo.Repo
	PwdRepo pwdRepo.Repo
	Rds     rdms.RandomString
	hmac    hmachash.HMAC
	pepper  string
}

// NewUserService will instantiate User Service
func NewUserService(
	repo userrepo.Repo,
	pwdRepo pwdRepo.Repo,
	rds rdms.RandomString,
	hmac hmachash.HMAC,
	pepper string) UserService {

	return &userService{
		Repo:    repo,
		PwdRepo: pwdRepo,
		Rds:     rds,
		hmac:    hmac,
		pepper:  pepper,
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

// Issue token for user to update his/her password
func (us *userService) InitiateResetPassowrd(email string) (string, error) {
	user, err := us.Repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	token, err := us.Rds.GenerateToken()
	if err != nil {
		return "", err
	}

	hashedToken := us.hmac.Hash(token)

	pwd := pwd.PasswordReset{
		UserID: user.ID,
		Token:  hashedToken,
	}

	err = us.PwdRepo.Create(&pwd)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *userService) CompleteUpdatePassword(token, newPassword string) (*user.User, error) {
	hashedToken := us.hmac.Hash(token)

	pwr, err := us.PwdRepo.GetOneByToken(hashedToken)
	if err != nil {
		return nil, err
	}

	// If the password rest is over 1 hours old, it is invalid
	if time.Now().Sub(pwr.CreatedAt) > (1 * time.Hour) {
		return nil, errors.New("Invalid Token")
	}

	user, err := us.Repo.GetByID(pwr.UserID)
	if err != nil {
		return nil, err
	}

	hashedPass, err := us.HashPassword(newPassword)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPass
	if err = us.Repo.Update(user); err != nil {
		return nil, err
	}

	if err = us.PwdRepo.Delete(pwr.ID); err != nil {
		fmt.Println("Failed to delete passwordreset record", pwr.ID, err.Error())
	}
	return user, nil
}
