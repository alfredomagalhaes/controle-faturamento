package controllers

import (
	"os"
	"time"

	"github.com/alfredomagalhaes/controle-faturamento/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {

	user := &models.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}

	err = user.Login()
	if err == models.ErrEmailNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "user email not found",
		})
	}
	if err == models.ErrInvalidCredentials {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "invalid credentials",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "connection error. please try again",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t,
		"expiration_time": claims["exp"]})

}
