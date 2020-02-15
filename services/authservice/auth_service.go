package authservice

import (
	"github.com/yhagio/go_api_boilerplate/domain/user"
	"time"

	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

// AuthService interface
type AuthService interface {
	IssueToken(u user.User) (string, error)
	ParseToken(token string) (*Claims, error)
}

type authService struct {
	jwtSecret string
}

// NewAuthService will instantiate AuthService
func NewAuthService(jwtSecret string) AuthService {
	return &authService{
		jwtSecret: jwtSecret,
	}
}

// Claims represents JWT claims
type Claims struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func (auth *authService) IssueToken(u user.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour) // 24 hours

	claims := Claims{
		u.Email,
		u.ID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Go API Boilerplate",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(auth.jwtSecret))
}

// ParseToken parsing token
func (auth *authService) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(auth.jwtSecret), nil
		},
	)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
