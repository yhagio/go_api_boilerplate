package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

type Claims struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	jwt.StandardClaims
}

// Remove "Bearer " from "Authorization" token string
func stripBearer(tok string) (string, error) {
	if len(tok) > 6 && strings.ToLower(tok[0:7]) == "bearer " {
		return tok[7:], nil
	}
	return tok, nil
}

// JWT is jwt middleware
func JWT(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := stripBearer(c.Request.Header.Get("Authorization"))
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		tokenClaims, err := jwt.ParseWithClaims(
			token,
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			},
		)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if tokenClaims != nil {
			claims, ok := tokenClaims.Claims.(*Claims)

			if ok && tokenClaims.Valid {
				c.Set("user_id", claims.ID)
				c.Set("user_email", claims.Email)

				c.Next()
				return
			}
		}

		c.AbortWithError(http.StatusUnauthorized, errors.New("NOP"))
	}
}
