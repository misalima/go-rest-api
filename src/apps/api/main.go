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
	if err != nil {
		panic("Erro ao conectar ao banco de dados: " + err.Error())
	}

	userRepository := postgres.NewUserRepository(pool)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	// Iniciar Echo
	e := echo.New()
	e.GET("/user/:id", userHandler.GetUserByID)
	e.GET("/users", userHandler.GetAllUsers)
	e.POST("/user", userHandler.CreateUser)
	e.PUT("/user/:id", userHandler.UpdateUser)
	e.DELETE("/user/:id", userHandler.DeleteUser)

	e.Logger.Fatal(e.Start(":8000"))

}
