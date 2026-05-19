package main

import (
	"os"
	"sysemp_feed/auth"
	"sysemp_feed/controller"
	"sysemp_feed/db"
	"sysemp_feed/repository"
	"sysemp_feed/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// =========================
	// DATABASE
	// =========================
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	defer dbConnection.Close()

	// =========================
	// JWT
	// =========================

	authService := auth.NewService(
		getEnv("JWT_SECRET", "dev-secret-change-me"),
		24*time.Hour,
	)

	// =========================
	// DEPENDENCY INJECTION
	// =========================
	baseRepository := repository.NewRepository(dbConnection)

	// User Configs
	UserCreateRepository := repository.NewUserRepository(baseRepository)
	UserUseCase := usecase.NewUserUseCase(UserCreateRepository)
	UserController := controller.NewUserController(UserUseCase)

	userRepository := repository.NewUserRepository(baseRepository)
	authUsecase := usecase.NewAuthUsecase(&userRepository)

	authController := controller.NewAuthController(
		authService,
		authUsecase,
	)

	// =========================
	// ROUTES
	// =========================
	server.POST("/login", authController.Login)
	server.POST("/create_user", UserController.CreateUser)
	server.DELETE("/approved_user/:id", UserController.ApproveUser)

	server.Run(":8080")
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
