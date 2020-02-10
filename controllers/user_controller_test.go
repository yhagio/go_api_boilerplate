package controllers

import (
	"encoding/json"
	"go_api_boilerplate/domain/user"
	"go_api_boilerplate/services/authservice"
	"go_api_boilerplate/services/userservice"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userRepo struct {
	mock.Mock
}

func (repo *userRepo) GetByID(id uint) (*user.User, error) {
	args := repo.Called(id)
	return args.Get(0).(*user.User), args.Error(1)
}

func (repo *userRepo) GetByEmail(email string) (*user.User, error) {
	args := repo.Called(email)
	return args.Get(0).(*user.User), args.Error(1)
}

func (repo *userRepo) Create(user *user.User) error {
	args := repo.Called(user)
	return args.Error(0)
}

func (repo *userRepo) Update(user *user.User) error {
	args := repo.Called(user)
	return args.Error(0)
}

// Output of HTTP Response Body structure
type output struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data user.User `json:"data"`
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestUserController(t *testing.T) {

	// Setup router + user controller
	ur := new(userRepo)
	us := userservice.NewUserService(ur, "pepper")
	as := authservice.NewAuthService("jwt-secret")
	userCtl := NewUserController(us, as)
	router := gin.Default()
	router.GET("/users/:id", userCtl.GetByID)

	t.Run("GetByID", func(t *testing.T) {
		// Test user
		user1 := &user.User{
			Email:     "alice@cc.cc",
			FirstName: "",
			LastName:  "",
			Active:    true,
			Role:      "standard",
		}
		// Stub UserService.GetByID func call
		ur.On("GetByID", uint(1)).Return(user1, nil)

		// Make HTTP Request to the testing endpoint
		w := performRequest(router, "GET", "/users/1")

		// Check statusCode
		assert.Equal(t, http.StatusOK, w.Code)

		// JSON to struct
		resBody := output{}
		json.NewDecoder(w.Body).Decode(&resBody)

		// Expected HTTP Response body structure
		expectedResBody := Response{
			Code: http.StatusOK,
			Msg:  "ok",
			Data: *user1,
		}

		assert.EqualValues(t, expectedResBody.Code, resBody.Code)
		assert.EqualValues(t, expectedResBody.Msg, resBody.Msg)
		assert.EqualValues(t, expectedResBody.Data, resBody.Data)
	})
}
