package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func SetConfig() {
	if os.Getenv("ENV") != "PROD" {
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error while reading config file %s", err)
		}
	} else {
		viper.AutomaticEnv()
	}
}
