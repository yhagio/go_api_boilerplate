package user_repo

import (
	"errors"
	"go_api_boilerplate/domain/user"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("postgres", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	gormDB.LogMode(true)
	defer gormDB.Close()

	t.Run("Get a user", func(t *testing.T) {
		expected := &user.User{
			Email: "alice@cc.cc",
		}

		u := NewUserRepo(gormDB)

		mock.
			ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."deleted_at" IS NULL AND \(\("users"\."id" \= 100\)\) ORDER BY "users"\."id" ASC LIMIT 1`).
			WillReturnRows(
				sqlmock.NewRows([]string{"email"}).
					AddRow("alice@cc.cc"))

		result, err := u.GetByID(100)

		assert.EqualValues(t, expected, result)
		assert.Nil(t, err)
	})

	t.Run("Error occurs", func(t *testing.T) {
		expected := errors.New("Nop")

		u := NewUserRepo(gormDB)

		mock.
			ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."deleted_at" IS NULL AND \(\("users"\."id" \= 100\)\) ORDER BY "users"\."id" ASC LIMIT 1`).
			WillReturnError(expected)

		result, err := u.GetByID(100)

		assert.EqualValues(t, expected, err)
		assert.Nil(t, result)
	})

	t.Run("Record Not Found", func(t *testing.T) {
		expected := errors.New("record not found")

		u := NewUserRepo(gormDB)

		mock.
			ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."deleted_at" IS NULL AND \(\("users"\."id" \= 100\)\) ORDER BY "users"\."id" ASC LIMIT 1`).
			WillReturnRows(
				sqlmock.NewRows([]string{}))

		result, err := u.GetByID(100)

		assert.EqualValues(t, expected, err)
		assert.Nil(t, result)
	})
}
