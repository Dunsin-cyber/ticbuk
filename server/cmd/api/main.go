package main

import (
	"fmt"

	"github.com/Dunsin-cyber/ticbuk/config"
	"github.com/Dunsin-cyber/ticbuk/db"
	"github.com/Dunsin-cyber/ticbuk/handlers"
	"github.com/Dunsin-cyber/ticbuk/middlewares"
	"github.com/Dunsin-cyber/ticbuk/repositories"
	"github.com/Dunsin-cyber/ticbuk/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	envConfig := config.NewEnvConfig()

	db := db.Init(envConfig, db.DBMigrator)

	app := echo.New()

	// Middleware
	app.Use(middleware.RequestLogger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allows your iPhone to connect
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// repository
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api/v1")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	//  privateRoute

	// handler
	handlers.NewEventHandler(server.Group("/events", middlewares.AuthProtected(db)), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket", middlewares.AuthProtected(db)), ticketRepository)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", envConfig.ServerPort)))

}
