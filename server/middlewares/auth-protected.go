package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Dunsin-cyber/ticbuk/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			ctx.Response().Header().Add("Vary", "Authorization")
			authHeader := ctx.Request().Header.Get("Authorization")

			if authHeader == "" {
				log.Warnf("missing authorization header")

				return ctx.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "fail",
					"message": "missing authorization header",
				})
			}
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				log.Warnf("invalid authorization header format")
				return ctx.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "fail",
					"message": "Unauthorized",
				})
			}

			tokenString := tokenParts[1]
			secret := []byte(os.Getenv("JWT_SECRET"))

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return secret, nil
			})
			if err != nil || !token.Valid {
				log.Warnf("invalid token:", err)
				return ctx.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "fail",
					"message": "Unauthorized",
				})
			}

			userId := token.Claims.(jwt.MapClaims)["id"]

			if err := db.Model(&models.User{}).Where("id = ?", userId).First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				log.Warnf("user not found:", err)
				return ctx.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "fail",
					"message": "Unauthorized",
				})
			}

			ctx.Set("userId", userId)
			return next(ctx)
		}
	}
}
