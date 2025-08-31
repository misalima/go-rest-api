package main

import (
	"os"
	"rest-api/src/apps/api/handlers"
	"rest-api/src/core/services"
	"rest-api/src/infra/repository/postgres"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Erro ao carregar .env: " + err.Error())
	}

	uri := os.Getenv("DATABASE_URL")

	pool, err := postgres.GetDBConnection(uri)

	userRepository := postgres.NewUserRepository(pool)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// Iniciar Echo
	e := echo.New()

	e.POST("/user", userHandler.CreateUser)
	e.GET("/user", userHandler.GetAllUsers)
	e.GET("/user/:id", userHandler.GetUserById)
	e.DELETE("/user/:id", userHandler.DeleteUserById)

	e.Logger.Fatal(e.Start(":8000"))

}
