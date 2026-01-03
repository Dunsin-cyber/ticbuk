package handlers

import (
	"context"
	"net/http"
	"strconv"
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
	router.GET("/:eventId", handler.GetOne)
	router.PUT("/:eventId", handler.UpdateOne)
	router.DELETE("/:eventId", handler.DeleteOne)

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
		"status":  "success",
		"message": "events retrieved successfully",
		"data":    events,
	})

}

func (h *EventHandler) GetOne(ctx echo.Context) error {
	eventId, _ := strconv.Atoi(ctx.Param("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	event, err := h.repository.GetOne(context, uint(eventId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "event retrieved successfully",
		"data":    event,
	})

}
func (h *EventHandler) CreateOne(ctx echo.Context) error {
	event := &models.Event{}
	if err := ctx.Bind(event); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	createdEvent, err := h.repository.CreateOne(context, event)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "event created successfully",
		"data":    createdEvent,
	})
}

func (h *EventHandler) UpdateOne(ctx echo.Context) error {
	eventId, _ := strconv.Atoi(ctx.Param("eventId"))
	updateData := make(map[string]interface{})

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.Echo().JSONSerializer.Deserialize(ctx, &updateData); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	event, err := h.repository.UpdateOne(context, uint(eventId), updateData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "event updated successfully",
		"data":    event,
	})

}

func (h *EventHandler) DeleteOne(ctx echo.Context) error {
	eventId, _ := strconv.Atoi(ctx.Param("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOne(context, uint(eventId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "event deleted successfully",
	})

}
