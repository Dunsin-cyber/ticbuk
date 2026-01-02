package main

import (
	"github.com/Dunsin-cyber/ticbuk/handlers"
	"github.com/Dunsin-cyber/ticbuk/repositories"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	// repository
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api/v1")

	// handler
	handlers.NewEventHandler(server.Group("/events"), eventRepository)

	app.Logger.Fatal(app.Start(":8000"))

}
