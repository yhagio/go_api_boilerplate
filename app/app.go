package app

import (
	"fmt"
	"go_api_boilerplate/configs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	router = gin.Default()
)

func init() {
	// Setup configs
	configs.SetConfig()

	// Setup services

	// Setup routes
	setRoutes()
}

func Run() {
	port := fmt.Sprintf(":%s", viper.Get("APP_PORT"))
	router.Run(port)
}
