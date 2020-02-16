package passwordreset

import (
	"database/sql/driver"
	"errors"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	pwd "github.com/yhagio/go_api_boilerplate/domain/passwordreset"
)

func setupDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("postgres", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	gormDB.LogMode(true)
	return gormDB, mock
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestGetOneByToken(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("Found a record", func(t *testing.T) {
		expected := &pwd.PasswordReset{
			Token: "token",
		}

		repo := NewPasswordResetRepo(gormDB)

		sqlStr := `SELECT * FROM "password_resets" WHERE "password_resets"."deleted_at" IS NULL AND ((token = $1)) ORDER BY "password_resets"."id" ASC LIMIT 1`

		mock.
			ExpectQuery(
				regexp.QuoteMeta(sqlStr)).
			WithArgs("token").
			WillReturnRows(
				sqlmock.NewRows([]string{"token"}).
					AddRow("token"))

		result, err := repo.GetOneByToken("token")

		assert.EqualValues(t, expected, result)
		assert.Nil(t, err)
	})

	t.Run("Record Not Found", func(t *testing.T) {
		expected := errors.New("record not found")

		repo := NewPasswordResetRepo(gormDB)

		sqlStr := `SELECT * FROM "password_resets" WHERE "password_resets"."deleted_at" IS NULL AND ((token = $1)) ORDER BY "password_resets"."id" ASC LIMIT 1`

		mock.
			ExpectQuery(regexp.QuoteMeta(sqlStr)).
			WithArgs("token").
			WillReturnRows(
				sqlmock.NewRows([]string{}))

		result, err := repo.GetOneByToken("token")

		assert.EqualValues(t, expected, err)
		assert.Nil(t, result)
	})
}

func TestCreate(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("Create a record", func(t *testing.T) {
		expected := &pwd.PasswordReset{
			UserID: uint(1),
			Token:  "token",
		}

		repo := NewPasswordResetRepo(gormDB)

		sqlStr := `INSERT INTO "password_resets" ("created_at","updated_at","deleted_at","user_id","token") VALUES ($1,$2,$3,$4,$5) RETURNING "password_resets"."id"`

		mock.ExpectBegin()

		mock.
			ExpectQuery(
				regexp.QuoteMeta(sqlStr)).
			WithArgs(AnyTime{}, AnyTime{}, nil, uint(1), "token").
			WillReturnRows(
				sqlmock.NewRows([]string{"token"}).
					AddRow(1))

		mock.ExpectCommit()

		err := repo.Create(expected)
		assert.Nil(t, err)
	})

	t.Run("Creating a record fails", func(t *testing.T) {
		exp := errors.New("oops")

		record := &pwd.PasswordReset{
			UserID: uint(1),
			Token:  "token",
		}
		repo := NewPasswordResetRepo(gormDB)

		sqlStr := `INSERT INTO "password_resets" ("created_at","updated_at","deleted_at","user_id","token") VALUES ($1,$2,$3,$4,$5) RETURNING "password_resets"."id"`

		mock.ExpectBegin()

		mock.
			ExpectQuery(
				regexp.QuoteMeta(sqlStr)).
			WithArgs(AnyTime{}, AnyTime{}, nil, uint(1), "token").
			WillReturnError(exp)

		mock.ExpectCommit()

		err := repo.Create(record)
		assert.NotNil(t, err)
		assert.EqualValues(t, exp, err)
	})
}

func TestDelete(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("Delete a record", func(t *testing.T) {
		repo := NewPasswordResetRepo(gormDB)

		mock.ExpectBegin()

		sqlStr := `UPDATE "password_resets" SET "deleted_at"=$1  WHERE "password_resets"."deleted_at" IS NULL AND "password_resets"."id" = $2`

		mock.
			ExpectExec(
				regexp.QuoteMeta(sqlStr)).
			WithArgs(AnyTime{}, uint(1)).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		err := repo.Delete(uint(1))
		assert.Nil(t, err)
	})

	t.Run("Deleting a record fails", func(t *testing.T) {
		exp := errors.New("oops")

		repo := NewPasswordResetRepo(gormDB)

		mock.ExpectBegin()

		sqlStr := `UPDATE "password_resets" SET "deleted_at"=$1  WHERE "password_resets"."deleted_at" IS NULL AND "password_resets"."id" = $2`

		mock.
			ExpectExec(
				regexp.QuoteMeta(sqlStr)).
			WithArgs(AnyTime{}, uint(1)).
			WillReturnError(exp)

		mock.ExpectCommit()

		err := repo.Delete(uint(1))
		assert.NotNil(t, err)
		assert.EqualValues(t, exp, err)
	})
}
