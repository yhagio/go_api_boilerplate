package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRoutes() {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Auth

	// Users

	// Items
}
