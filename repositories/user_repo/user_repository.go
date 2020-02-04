package user_repo

import (
	"go_api_boilerplate/domain/user"

	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	GetById(id int) (*user.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetById(id int) (*user.User, error) {
	var user user.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
