package main

import (
	"sysemp_feed/controller"
	"sysemp_feed/db"
	"sysemp_feed/repository"
	"sysemp_feed/usecase"

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
	// DEPENDENCY INJECTION
	// =========================
	baseRepository := repository.NewRepository(dbConnection)

	// User Configs
	UserCreateRepository := repository.NewUserRepository(baseRepository)
	UserUseCase := usecase.NewUserUseCase(UserCreateRepository)
	UserController := controller.NewUserController(UserUseCase)

	// =========================
	// ROUTES
	// =========================
	server.POST("/user", UserController.CreateUser)

	server.Run(":8080")

}
