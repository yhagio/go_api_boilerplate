package configs

import (
	"os"
)

// MailgunConfig object
type MailgunConfig struct {
	APIKey       string `env:"MAILGUN_API_KEY"`
	// PublicAPIKey string `env:"MAILGUN_PUBLIC_KEY"`
	Domain       string `env:"MAILGUN_DOMAIN"`
}

// GetMailgunConfig get Mainlgun config object
func GetMailgunConfig() MailgunConfig {
	return MailgunConfig{
		APIKey:       os.Getenv("MAILGUN_API_KEY"),
		// PublicAPIKey: os.Getenv("MAILGUN_PUBLIC_KEY"),
		Domain:       os.Getenv("MAILGUN_DOMAIN"),
	}
}
