package controllers

import (
	"os"
	"strings"

	b64 "encoding/base64"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	"github.com/alfredomagalhaes/controle-faturamento/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

//Create method to create user in the database base and return authenticatin=on token
func CreateUser(c *fiber.Ctx) error {

	var userEntry models.User

	h := c.Request().Header.Peek("Authorization")

	if !strings.Contains(string(h), "Basic") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "usuário não autorizado.",
		})
	}
	userRequest := strings.Split(string(h), " ")[1]
	userDecript, _ := b64.StdEncoding.DecodeString(userRequest)
	if string(userDecript) != os.Getenv("ADMIN_USER")+":"+os.Getenv("ADMIN_PASS") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "usuário não autorizado.",
		})
	}

	err := c.BodyParser(&userEntry)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Dados do usuário incorreto, verificar informações",
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userEntry.Password), bcrypt.DefaultCost)
	userEntry.Password = string(hashedPassword)

	if config.MI.DB.Create(&userEntry).Error != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Erro ao criar o usuário.",
		})
	}

	//Create new JWT token for the newly registered User
	tk := &models.Token{UserID: userEntry.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	userEntry.Token = tokenString

	userEntry.Password = "" //delete password

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"user":    userEntry,
	})
}
