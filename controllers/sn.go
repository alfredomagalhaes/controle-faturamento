package controllers

import (
	"github.com/alfredomagalhaes/controle-faturamento/models"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func CriarTabelaSN(c *fiber.Ctx) error {

	tb := &models.SN_Tabela{}

	err := c.BodyParser(&tb)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
			"error":   err,
		})
	}

	err = tb.CriarTabelaSN()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "",
		"data":    tb,
	})
}

func ObterTabelaSN(c *fiber.Ctx) error {

	//tb := &models.SN_Tabela{}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}
	idPesq, _ := uuid.FromString(id)
	tb, err := models.ObterTabelaSN(idPesq)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "nenhum registro localizado com o ID informado",
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "",
		"data":    tb,
	})
}
