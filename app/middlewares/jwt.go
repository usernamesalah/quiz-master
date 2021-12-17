package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/usernamesalah/quiz-master/internal/config"
	"github.com/usernamesalah/quiz-master/internal/constants"
)

func AuthenticationMiddleware(g *echo.Group, user constants.User) {
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.AppConfig.JWTSecret),
	}))

	switch user {
	case constants.Admin:
		{
			g.Use(validateJWTAdmin)
			break
		}
	case constants.Client:
		{
			g.Use(validateJWTclient)
			break
		}
	default:
		{
			g.Use(validateJWTclient)
			break
		}
	}
}

func validateJWTclient(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if fmt.Sprintf("%s", claims["uid"]) != "0" || fmt.Sprintf("%s", claims["uid"]) != "" {
				next(c)
			}
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Please Sign In \n Woops! Gonna sign in first\n Only a click away and you can continue to enjoy")
	}
}

func validateJWTAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["isAdmin"].(bool) && fmt.Sprintf("%s", claims["uid"]) != "0" {
				next(c)
			}
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Please Sign In \n Woops! Gonna sign in first\n Only a click away and you can continue to enjoy")
	}
}
