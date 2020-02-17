package middlewares

import (
	"net/http"
	"net/http/httptest"

	"github.com/yhagio/go_api_boilerplate/domain/user"
	"github.com/yhagio/go_api_boilerplate/services/authservice"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequireLoggedInMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	alice := &user.User{
		Email:     "alice@cc.cc",
		FirstName: "",
		LastName:  "",
		Active:    false,
		Role:      "",
	}

	svc := authservice.NewAuthService("secret")

	t.Run("Has valid token and authorized", func(t *testing.T) {
		token, _ := svc.IssueToken(*alice)
		bearerToken := "Bearer " + token

		router.GET("/test1", RequireLoggedIn("secret"), func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			assert.EqualValues(t, 0, userID)

			email, _ := c.Get("user_email")
			assert.EqualValues(t, "alice@cc.cc", email)

			c.Status(http.StatusOK)
		})

		request, _ := http.NewRequest("GET", "/test1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		request.Header.Add("Authorization", bearerToken)
		c.Request = request
		router.ServeHTTP(w, request)

		assert.EqualValues(t, http.StatusOK, w.Code)
	})

	t.Run("Unauthorized without token", func(t *testing.T) {
		router.GET("/test2", RequireLoggedIn("secret"), func(c *gin.Context) {
			email, _ := c.Get("user_email")
			assert.Nil(t, email)

			userID, _ := c.Get("user_id")
			assert.Nil(t, userID)

			c.Status(http.StatusOK)
		})

		request, _ := http.NewRequest("GET", "/test2", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = request
		router.ServeHTTP(w, request)

		assert.EqualValues(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Unauthorized without valid token", func(t *testing.T) {
		router.GET("/test3", RequireLoggedIn("secret"), func(c *gin.Context) {
			email, _ := c.Get("user_email")
			assert.Nil(t, email)

			userID, _ := c.Get("user_id")
			assert.Nil(t, userID)

			c.Status(http.StatusOK)
		})

		request, _ := http.NewRequest("GET", "/test3", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		request.Header.Add("Authorization", "Bearer token")
		c.Request = request
		router.ServeHTTP(w, request)

		assert.EqualValues(t, http.StatusUnauthorized, w.Code)
	})
}
