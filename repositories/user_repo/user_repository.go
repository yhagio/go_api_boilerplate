package user_repo

import (
	"go_api_boilerplate/domain/user"

	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	GetByID(id uint) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	Create(user *user.User) error
	Update(user *user.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetByID(id uint) (*user.User, error) {
	var user user.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetByEmail(email string) (*user.User, error) {
	var user user.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) Create(user *user.User) error {
	return u.db.Create(user).Error
}

func (ug *userRepo) Update(user *user.User) error {
	return ug.db.Save(user).Error
}
