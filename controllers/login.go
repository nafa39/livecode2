package controllers

import (
	"net/http"

	models "livecode2/models"
	"livecode2/utils"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	var request models.LoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	user, err := models.GetUserByEmail(request.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	// Compare password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid password"})
	}

	// Generate JWT Token
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
