package main

import (
	"fmt"

	"github.com/Dunsin-cyber/ticbuk/config"
	"github.com/Dunsin-cyber/ticbuk/db"
	"github.com/Dunsin-cyber/ticbuk/handlers"
	"github.com/Dunsin-cyber/ticbuk/repositories"
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

	// repository
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	// Routing
	server := app.Group("/api/v1")

	// handler
	handlers.NewEventHandler(server.Group("/events"), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", envConfig.ServerPort)))

}
