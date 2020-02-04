package app

import (
	"fmt"
	"go_api_boilerplate/configs"
	"go_api_boilerplate/domain/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"go_api_boilerplate/controllers"
	"go_api_boilerplate/repositories/user_repo"
	"go_api_boilerplate/services/user_service"

	_ "github.com/lib/pq" // For Postgres setup
)

var (
	router = gin.Default()
)

// Run application
func Run() {
	// ====== Setup configs ============
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()

	// Connects to PostgresDB
	db, err := gorm.Open(
		config.Postgres.Dialect(),
		config.Postgres.GetPostgresConnectionInfo(),
	)
	if err != nil {
		panic(err)
	}

	// Migration
	db.AutoMigrate(&user.User{})
	defer db.Close()

	// ====== Setup infra ==============

	// ====== Setup repositories =======
	userRepo := user_repo.NewUserRepo(db)

	// ====== Setup services ===========
	userService := user_service.NewUserService(userRepo)

	// ====== Setup controllers ========
	userController := controllers.NewUserController(userService)

	// ====== Setup routes =============
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/users/:id", userController.GetById)

	// Run
	// port := fmt.Sprintf(":%s", viper.Get("APP_PORT"))
	port := fmt.Sprintf(":%s", config.Port)
	router.Run(port)
}
