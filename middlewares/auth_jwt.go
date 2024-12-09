// package middleware

// import (
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/labstack/echo/middleware"
// 	"github.com/labstack/echo/v4"
// )

// var JwtKey = middleware.JWTWithConfig(middleware.JWTConfig{
// 	SigningKey:  []byte("secret"),
// 	Claims:      &jwt.StandardClaims{},
// 	TokenLookup: "header:Authorization",
// 	AuthScheme:  "Bearer",
// 	ErrorHandler: func(err error) error {
// 		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authentication", err.Error())
// 	},
// })

package middlewares

import (
	"livecode2/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetJWTSecret()), nil
		})
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		return next(c)
	}
}
