package passwordreset

import (
	"github.com/jinzhu/gorm"
	pwd "github.com/yhagio/go_api_boilerplate/domain/passwordreset"
)

// Repo interface
type Repo interface {
	GetOneByToken(token string) (*pwd.PasswordReset, error)
	Create(pr *pwd.PasswordReset) error
	Delete(id uint) error
}

type pwdRepo struct {
	db *gorm.DB
}

// NewPasswordResetRepo will instantiate User Repository
func NewPasswordResetRepo(db *gorm.DB) Repo {
	return &pwdRepo{
		db: db,
	}
}

func (repo *pwdRepo) GetOneByToken(token string) (*pwd.PasswordReset, error) {
	var pwr pwd.PasswordReset
	if err := repo.db.Where("token = ?", token).First(&pwr).Error; err != nil {
		return nil, err
	}
	return &pwr, nil
}

func (repo *pwdRepo) Create(pwr *pwd.PasswordReset) error {
	return repo.db.Create(pwr).Error
}

func (repo *pwdRepo) Delete(id uint) error {
	pwr := pwd.PasswordReset{
		Model: gorm.Model{ID: id},
	}
	return repo.db.Delete(&pwr).Error
}
