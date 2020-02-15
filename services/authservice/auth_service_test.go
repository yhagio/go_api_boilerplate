package authservice

import (
	"github.com/yhagio/go_api_boilerplate/domain/user"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	t.Run("Generate token", func(t *testing.T) {
		u := &user.User{
			gorm.Model{ID: uint(1), CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil},
			"",
			"",
			"alice@cc.cc",
			"",
			"",
			true,
		}

		svc := NewAuthService("secret")

		token, err := svc.IssueToken(*u)
		assert.Nil(t, err)
		assert.IsType(t, "string", token)
	})

	t.Run("Invalid Token 1", func(t *testing.T) {
		token := "hello"
		svc := NewAuthService("secret")
		claims, err := svc.ParseToken(token)
		assert.NotNil(t, err)
		assert.Nil(t, claims)
	})

	t.Run("Invalid Token 2", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFsaWNlQGNjLmNjIiwiaWQiOjIsImV4cCI6MTU4MTM1ODgyMywiaXNzIjoiR28gQVBJIEJvaWxlcnBsYXRlIn0.cB2G3sH6Mlyi0xHVKVH1QT4UGmv-Co36C0mbgIoQc80"
		svc := NewAuthService("secret")
		claims, err := svc.ParseToken(token)
		assert.NotNil(t, err)
		assert.Nil(t, claims)
	})

	t.Run("Generate and validate token", func(t *testing.T) {
		u := &user.User{
			gorm.Model{ID: uint(1), CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil},
			"",
			"",
			"alice@cc.cc",
			"",
			"",
			true,
		}

		svc := NewAuthService("secret")

		token, err := svc.IssueToken(*u)
		assert.Nil(t, err)
		assert.IsType(t, "string", token)

		claims, err := svc.ParseToken(token)
		assert.Nil(t, err)
		assert.NotNil(t, claims)

		assert.EqualValues(t, u.ID, claims.ID)
		assert.EqualValues(t, u.Email, claims.Email)
	})
}
