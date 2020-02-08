package auth_service

import (
	"crypto/md5"
	"encoding/hex"
	"go_api_boilerplate/domain/user"
	"time"

	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

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

type Claims struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func (auth *authService) IssueToken(u user.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		u.Email,
		u.ID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Go API Boilerplate",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(auth.jwtSecret))

	return token, err
}

// ParseToken parsing token
func (auth *authService) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return auth.jwtSecret, nil
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

func (auth *authService) EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
