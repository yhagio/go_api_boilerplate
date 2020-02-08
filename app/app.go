package app

import (
	"fmt"
	"go_api_boilerplate/configs"
	"go_api_boilerplate/domain/user"
	"go_api_boilerplate/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"go_api_boilerplate/controllers"
	"go_api_boilerplate/repositories/user_repo"
	"go_api_boilerplate/services/auth_service"
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
	// db.DropTableIfExists(&user.User{})
	db.AutoMigrate(&user.User{})
	defer db.Close()

	// ====== Setup infra ==============

	// ====== Setup repositories =======
	userRepo := user_repo.NewUserRepo(db)

	// ====== Setup services ===========
	userService := user_service.NewUserService(userRepo, config.Pepper)
	authService := auth_service.NewAuthService(config.SigningKey)

	// ====== Setup controllers ========
	userCtl := controllers.NewUserController(userService, authService)

	// ====== Setup middlewares ========
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// ====== Setup routes =============
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	api := router.Group("/api")

	api.POST("/register", userCtl.Register)
	api.POST("/login", userCtl.Login)
	api.POST("/forgot_password", userCtl.ForgotPassword)
	api.POST("/reset_password", userCtl.ResetPassword)

	user := api.Group("/users")

	user.GET("/:id", userCtl.GetByID)

	account := api.Group("/account")
	account.Use(middlewares.JWT(config.SigningKey))
	{
		account.GET("/profile", userCtl.GetProfile)
		account.PUT("/profile", userCtl.Update)
	}

	// Run
	// port := fmt.Sprintf(":%s", viper.Get("APP_PORT"))
	port := fmt.Sprintf(":%s", config.Port)
	router.Run(port)
}
