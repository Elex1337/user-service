package main

import (
	"fmt"
	"github.com/Elex1337/user-service/api/handler"
	"github.com/Elex1337/user-service/config"
	_ "github.com/Elex1337/user-service/docs"
	"github.com/Elex1337/user-service/internal/db"
	"github.com/Elex1337/user-service/internal/repository"
	"github.com/Elex1337/user-service/internal/service"
	route "github.com/Elex1337/user-service/routes"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"os"
)

// @title User Service API
// @description User API
// @host localhost:8080
func main() {
	dbConfig := config.LoadConfig()

	database := db.ConnectDB(dbConfig)
	defer database.Close()
	db.RunMigrations(database, "./migrations")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	e := echo.New()

	route.Routes(e, userHandler)

	e.Start(fmt.Sprintf(":%s", port))
}
