package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

var JwtKey = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey:  []byte("secret"),
	Claims:      &jwt.StandardClaims{},
	TokenLookup: "header:Authorization",
	AuthScheme:  "Bearer",
	ErrorHandler: func(err error) error {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authentication", err.Error())
	},
})
