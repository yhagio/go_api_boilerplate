package configs

import "os"

const (
	prod = "production"
)

// Config object
type Config struct {
	Env       string         `env:"ENV"`
	Pepper    string         `env:"PEPPER"`
	HMACKey   string         `env:"HMAC_KEY"`
	Postgres  PostgresConfig `json:"postgres"`
	Mailgun   MailgunConfig  `json:"mailgun"`
	JWTSecret string         `env:"JWT_SIGN_KEY"`
	Host      string         `env:"APP_HOST"`
	Port      string         `env:"APP_PORT"`
	FromEmail string         `env:"EMAIL_FROM"`
}

// IsProd Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == prod
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:       os.Getenv("ENV"),
		Pepper:    os.Getenv("PEPPER"),
		HMACKey:   os.Getenv("HMAC_KEY"),
		Postgres:  GetPostgresConfig(),
		Mailgun:   GetMailgunConfig(),
		JWTSecret: os.Getenv("JWT_SIGN_KEY"),
		Host:      os.Getenv("APP_HOST"),
		Port:      os.Getenv("APP_PORT"),
		FromEmail: os.Getenv("EMAIL_FROM"),
	}
}

//------------- Viper Implementation -------------------

// import (
// 	"log"
// 	"os"

// 	"github.com/spf13/viper"
// )

// func SetConfig() {
// 	if os.Getenv("ENV") != "PROD" {
// 		viper.SetConfigFile(".env")

// 		if err := viper.ReadInConfig(); err != nil {
// 			log.Fatalf("Error while reading config file %s", err)
// 		}
// 	} else {
// 		viper.AutomaticEnv()
// 	}
// }
