package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Dunsin-cyber/ticbuk/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type AuthHandler struct {
	authService models.AuthService
}

func NewAuthHandler(router *echo.Group, service models.AuthService) {
	handler := &AuthHandler{
		authService: service,
	}

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	creds := &models.AuthCredentials{}
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.Bind(&creds); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err := validate.Struct(creds); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	token, user, err := h.authService.Login(context, creds)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "user logged in successfully",
		"data": map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})

}

func (h *AuthHandler) Register(ctx echo.Context) error {
	creds := &models.AuthCredentials{}
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.Bind(&creds); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err := validate.Struct(creds); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": fmt.Errorf("please provide a valid email and password").Error(),
		})
	}

	token, user, err := h.authService.Register(context, creds)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "user logged in successfully",
		"data": map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})

}
