package userhandler

import (
	"livecode2/config"
	internal "livecode2/internal/userdto"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {
	var reg internal.RegisterRequest
	if err := c.Bind(&reg); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error(), "message": "Invalid request"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error(), "message": "Invalid generate password"})
	}

	user := internal.User{
		Email:    reg.Email,
		Password: string(hashedPassword),
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error(), "message": "Register failed"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Register success"})
}
