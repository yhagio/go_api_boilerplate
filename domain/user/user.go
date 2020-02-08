package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	// Username  string `gorm:"NOT NULL; size:255; UNIQUE_INDEX"`
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Email     string `gorm:"NOT NULL; UNIQUE_INDEX"`
	// Password     string `gorm:"-"`
	Password string `gorm:"NOT NULL"`
	// Token        string `gorm:"-"`
	// Token    string `gorm:"NOT NULL; UNIQUE_INDEX"`
	// Admin        bool   `gorm:"NOT NULL; DEFAULT: false"`
	Role   string `gorm:"NOT_NULL;size:255;DEFAULT:'standard'"`
	Active bool   `gorm:"NOT NULL; DEFAULT: true"`
}
