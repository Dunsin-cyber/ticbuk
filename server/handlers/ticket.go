package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Dunsin-cyber/ticbuk/models"
	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	repository models.TicketRepository
}

func NewTicketHandler(router *echo.Group, repository models.TicketRepository) {
	handler := &TicketHandler{
		repository: repository,
	}

	router.GET("", handler.GetMany)
	router.POST("", handler.CreateOne)
	router.GET("/:ticketId", handler.GetOne)
	router.PUT("/validate", handler.ValidateOne)

}

func (h *TicketHandler) GetMany(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	tickets, err := h.repository.GetMany(context)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "tickets retrieved successfully",
		"data":    tickets,
	})

}

func (h *TicketHandler) GetOne(ctx echo.Context) error {
	ticketId, _ := strconv.Atoi(ctx.Param("ticketId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	ticket, err := h.repository.GetOne(context, uint(ticketId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "ticket retrieved successfully",
		"data":    ticket,
	})

}
func (h *TicketHandler) CreateOne(ctx echo.Context) error {
	ticket := &models.Ticket{}
	if err := ctx.Bind(ticket); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	createdTicket, err := h.repository.CreateOne(context, ticket)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "ticket created successfully",
		"data":    createdTicket,
	})
}

func (h *TicketHandler) ValidateOne(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	validateBody := &models.ValidateTicket{}

	if err := ctx.Bind(validateBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true

	ticket, err := h.repository.UpdateOne(context, validateBody.TicketID, validateData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "welcome to the show!",
		"data":    ticket,
	})
}

