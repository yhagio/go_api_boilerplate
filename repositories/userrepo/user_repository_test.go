package userrepo

import (
	"database/sql/driver"
	"errors"
	"github.com/yhagio/go_api_boilerplate/domain/user"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
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

func TestGetByID(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("Get a user", func(t *testing.T) {
		expected := &user.User{
			Email: "alice@cc.cc",
		}

		u := NewUserRepo(gormDB)

		mock.
			ExpectQuery(
				regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND (("users"."id" = 100)) ORDER BY "users"."id" ASC LIMIT 1`)).
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
			ExpectQuery(
				regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND (("users"."id" = 100)) ORDER BY "users"."id" ASC LIMIT 1`)).
			WillReturnError(expected)

		result, err := u.GetByID(100)

		assert.EqualValues(t, expected, err)
		assert.Nil(t, result)
	})

	t.Run("Record Not Found", func(t *testing.T) {
		expected := errors.New("record not found")

		u := NewUserRepo(gormDB)

		mock.
			ExpectQuery(
				regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND (("users"."id" = 100)) ORDER BY "users"."id" ASC LIMIT 1`)).
			WillReturnRows(
				sqlmock.NewRows([]string{}))

		result, err := u.GetByID(100)

		assert.EqualValues(t, expected, err)
		assert.Nil(t, result)
	})
}

func TestGetByEmail(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("Get a user", func(t *testing.T) {
		expected := &user.User{
			Email: "alice@cc.cc",
		}

		u := NewUserRepo(gormDB)
		sqlStr := `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((email = $1)) ORDER BY "users"."id" ASC LIMIT 1`

		mock.
			ExpectQuery(regexp.QuoteMeta(sqlStr)).
			WithArgs("alice@cc.cc").
			WillReturnRows(
				sqlmock.NewRows([]string{"email"}).
					AddRow("alice@cc.cc"))

		result, err := u.GetByEmail("alice@cc.cc")

		assert.EqualValues(t, expected, result)
		assert.Nil(t, err)
	})

	t.Run("Error occurs", func(t *testing.T) {
		expected := errors.New("Nop")

		u := NewUserRepo(gormDB)
		sqlStr := `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((email = $1)) ORDER BY "users"."id" ASC LIMIT 1`

		mock.
			ExpectQuery(regexp.QuoteMeta(sqlStr)).
			WithArgs("alice@cc.cc").
			WillReturnError(expected)

		result, err := u.GetByEmail("alice@cc.cc")

		assert.EqualValues(t, expected, err)
		assert.Nil(t, result)
	})

	t.Run("Record Not Found", func(t *testing.T) {
		expected := errors.New("record not found")

		u := NewUserRepo(gormDB)
		sqlStr := `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((email = $1)) ORDER BY "users"."id" ASC LIMIT 1`

		mock.
			ExpectQuery(regexp.QuoteMeta(sqlStr)).
			WithArgs("alice@cc.cc").
			WillReturnRows(
				sqlmock.NewRows([]string{}))

		result, err := u.GetByEmail("alice@cc.cc")

		assert.EqualValues(t, expected, err)
		assert.Nil(t, result)
	})
}

func TestCreate(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("Create a user", func(t *testing.T) {
		user := &user.User{
			Email:    "alice@cc.cc",
			Password: "abc",
		}

		u := NewUserRepo(gormDB)

		mock.ExpectBegin()

		mock.
			ExpectQuery(
				regexp.QuoteMeta(
					`INSERT INTO "users" ("created_at","updated_at","deleted_at","first_name","last_name","email","password") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "users"."id"`)).
			WithArgs(AnyTime{}, AnyTime{}, nil, "", "", "alice@cc.cc", "abc").
			WillReturnRows(
				sqlmock.NewRows([]string{"id"}).
					AddRow(1))

		mock.ExpectCommit()

		err := u.Create(user)
		assert.Nil(t, err)
	})

	t.Run("Create a user fails", func(t *testing.T) {
		exp := errors.New("oops")
		user := &user.User{
			Email:    "alice@cc.cc",
			Password: "abc",
		}

		u := NewUserRepo(gormDB)

		mock.ExpectBegin()

		mock.
			ExpectQuery(
				regexp.QuoteMeta(
					`INSERT INTO "users" ("created_at","updated_at","deleted_at","first_name","last_name","email","password") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "users"."id"`)).
			WithArgs(AnyTime{}, AnyTime{}, nil, "", "", "alice@cc.cc", "abc").
			WillReturnError(exp)

		mock.ExpectCommit()

		err := u.Create(user)
		assert.NotNil(t, err)
		assert.EqualValues(t, exp, err)
	})
}

func TestUpdate(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("Update a user", func(t *testing.T) {
		user := &user.User{
			Email:    "alice@cc.cc",
			Password: "abc",
		}

		u := NewUserRepo(gormDB)

		mock.ExpectBegin()

		mock.
			ExpectQuery(
				regexp.QuoteMeta(
					`INSERT INTO "users" ("created_at","updated_at","deleted_at","first_name","last_name","email","password") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "users"."id"`)).
			WithArgs(AnyTime{}, AnyTime{}, nil, "", "", "alice@cc.cc", "abc").
			WillReturnRows(
				sqlmock.NewRows([]string{"id"}).
					AddRow(1))

		mock.ExpectCommit()

		err := u.Update(user)
		assert.Nil(t, err)
	})

	t.Run("Update a user fails", func(t *testing.T) {
		exp := errors.New("oops")
		user := &user.User{
			Email:    "alice@cc.cc",
			Password: "abc",
		}

		u := NewUserRepo(gormDB)

		mock.ExpectBegin()

		mock.
			ExpectQuery(
				regexp.QuoteMeta(
					`INSERT INTO "users" ("created_at","updated_at","deleted_at","first_name","last_name","email","password") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "users"."id"`)).
			WithArgs(AnyTime{}, AnyTime{}, nil, "", "", "alice@cc.cc", "abc").
			WillReturnError(exp)

		mock.ExpectCommit()

		err := u.Update(user)
		assert.NotNil(t, err)
		assert.EqualValues(t, exp, err)
	})
}
