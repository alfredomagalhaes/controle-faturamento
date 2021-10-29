package controllers

import (
	"github.com/alfredomagalhaes/controle-faturamento/models"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func CriarTabelaINSS(c *fiber.Ctx) error {

	tb := &models.INSS_Tabela{}

	err := c.BodyParser(&tb)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	err = tb.CriarTabelaINSS()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "",
		"data":    tb,
	})
}

func AtualizarTabelaINSS(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id da tabela não informado",
		})
	}

	tb := &models.INSS_Tabela{}

	err := c.BodyParser(&tb)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	tb.ID, _ = uuid.FromString(id)
	err = tb.AtualizarTabelaINSS()

	if err != nil {
		return c.Status(fiber.StatusPreconditionFailed).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "",
		"data":    tb,
	})
}

func ObterTabelaINSS(c *fiber.Ctx) error {

	//tb := &models.SN_Tabela{}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}
	idPesq, _ := uuid.FromString(id)
	tb, err := models.ObterTabelaINSS(idPesq)
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

func ObterFaixasPorTabelaINSS(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}
	idPesq, _ := uuid.FromString(id)
	tb, err := models.ObterFaixasPorTabelaIR(idPesq)
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

func ApagarFaixaINSS(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}
	idPesq, _ := uuid.FromString(id)
	err := models.ApagarFaixasIR(idPesq)
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
	})
}
