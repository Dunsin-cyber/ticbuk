package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Dunsin-cyber/ticbuk/models"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	repository models.EventRepository
}

func NewEventHandler(router *echo.Group, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	router.GET("", handler.GetMany)
	router.POST("", handler.CreateOne)
	router.GET(":eventId", handler.GetOne)

}

func (h *EventHandler) GetMany(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	events, err := h.repository.GetMany(context)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   events,
	})

}

func (h *EventHandler) GetOne(ctx echo.Context) error {
	return nil
}
func (h *EventHandler) CreateOne(ctx echo.Context) error {
	return nil
}
