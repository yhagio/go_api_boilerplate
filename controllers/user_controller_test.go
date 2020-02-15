package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yhagio/go_api_boilerplate/domain/user"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// NOTE: Mocked services are in './user_controller_setup_test.go'

// Output of HTTP Response Body structure
type output struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data user.User `json:"data"`
}

type outputAuth struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data gin.H  `json:"data"`
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestUserController(t *testing.T) {

	// Setup router + user controller
	us := &userSvc{}
	as := &authSvc{"jwt-secret"}
	es := &emailSvc{}
	userCtl := NewUserController(us, as, es)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users/:id", userCtl.GetByID)
	// router.GET("/profile", userCtl.GetProfile)
	// router.POST("/register", userCtl.Register)
	// router.POST("/login", userCtl.Login)
	// router.PU("/profile", userCtl.Update)

	// Using router version
	t.Run("GetByID", func(t *testing.T) {
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
			Data: *sampleUser,
		}

		assert.EqualValues(t, expectedResBody.Code, resBody.Code)
		assert.EqualValues(t, expectedResBody.Msg, resBody.Msg)
		assert.EqualValues(t, expectedResBody.Data, resBody.Data)
	})

	// Without using router version
	t.Run("GetProfile", func(t *testing.T) {
		// Mock HTTP Request to the testing endpoint
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user_id", uint(0))

		userCtl.GetProfile(c)

		// Check statusCode
		assert.Equal(t, http.StatusOK, w.Code)

		// JSON to struct
		resBody := output{}
		json.NewDecoder(w.Body).Decode(&resBody)

		// Expected HTTP Response body structure
		expectedResBody := Response{
			Code: http.StatusOK,
			Msg:  "ok",
			Data: *sampleUser,
		}

		assert.EqualValues(t, expectedResBody.Code, resBody.Code)
		assert.EqualValues(t, expectedResBody.Msg, resBody.Msg)
		assert.EqualValues(t, expectedResBody.Data, resBody.Data)
	})

	t.Run("Register", func(t *testing.T) {
		reqBody := map[string]string{
			"email":    "alice@cc.cc",
			"password": "123test",
		}

		// Mock HTTP Request to the testing endpoint
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Mock Request body
		payload, _ := json.Marshal(reqBody)
		request := httptest.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		// request.Header.Set("content-type", "application/json")
		// router.ServeHTTP(w, request)
		c.Request = request

		userCtl.Register(c)

		// Check statusCode
		assert.Equal(t, http.StatusOK, w.Code)

		// Response body JSON to struct
		resBody := Response{}
		json.NewDecoder(w.Body).Decode(&resBody)

		// Expected HTTP Response body structure
		expectedResBody := Response{
			Code: http.StatusOK,
			Msg:  "ok",
			Data: map[string]interface{}{
				"token": "nice-token",
				"user": map[string]interface{}{
					"id":        float64(0),
					"email":     "alice@cc.cc",
					"firstName": "",
					"lastName":  "",
					"role":      "",
					"active":    false,
				},
			},
		}

		assert.EqualValues(t, expectedResBody, resBody)
	})

	t.Run("Login", func(t *testing.T) {
		reqBody := map[string]string{
			"email":    "alice@cc.cc",
			"password": "123test",
		}

		// Mock HTTP Request to the testing endpoint
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Mock Request body
		payload, _ := json.Marshal(reqBody)
		request := httptest.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		// request.Header.Set("content-type", "application/json")
		// router.ServeHTTP(w, request)
		c.Request = request

		userCtl.Login(c)

		// Check statusCode
		assert.Equal(t, http.StatusOK, w.Code)

		// Response body JSON to struct
		resBody := Response{}
		json.NewDecoder(w.Body).Decode(&resBody)

		// Expected HTTP Response body structure
		expectedResBody := Response{
			Code: http.StatusOK,
			Msg:  "ok",
			Data: map[string]interface{}{
				"token": "nice-token",
				"user": map[string]interface{}{
					"id":        float64(0),
					"email":     "alice@cc.cc",
					"firstName": "",
					"lastName":  "",
					"role":      "",
					"active":    false,
				},
			},
		}

		assert.EqualValues(t, expectedResBody, resBody)
	})

	t.Run("Update", func(t *testing.T) {
		reqBody := map[string]string{
			"email":     "alice@cc.cc",
			"firstName": "alice",
			"lastName":  "smith",
		}

		// Mock HTTP Request to the testing endpoint
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("user_id", uint(0))

		// Mock Request body
		payload, _ := json.Marshal(reqBody)
		request := httptest.NewRequest("PUT", "/profile", bytes.NewBuffer(payload))
		// request.Header.Set("content-type", "application/json")
		// router.ServeHTTP(w, request)
		c.Request = request

		userCtl.Update(c)

		// Check statusCode
		assert.Equal(t, http.StatusOK, w.Code)

		// Response body JSON to struct
		resBody := Response{}
		json.NewDecoder(w.Body).Decode(&resBody)

		// Expected HTTP Response body structure
		expectedResBody := Response{
			Code: http.StatusOK,
			Msg:  "ok",
			Data: map[string]interface{}{
				"id":        float64(0),
				"email":     "alice@cc.cc",
				"firstName": "alice",
				"lastName":  "smith",
				"role":      "",
				"active":    false,
			},
		}

		assert.EqualValues(t, expectedResBody, resBody)
	})
}
