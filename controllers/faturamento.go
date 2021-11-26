package controllers

import (
	"strconv"

	"github.com/alfredomagalhaes/controle-faturamento/models"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func InserirFaturamento(c *fiber.Ctx) error {

	tb := &models.Faturamento{}

	err := c.BodyParser(&tb)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	err = tb.InserirFaturamento()
	if err != nil {
		if err == models.ErrReferenciaJaCadastrada {
			return c.Status(fiber.StatusFound).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "",
		"data":    tb,
	})
}

func ObterFaturamento(c *fiber.Ctx) error {

	//tb := &models.SN_Tabela{}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}
	idPesq, _ := uuid.FromString(id)
	tb, err := models.ObterFaturamento(idPesq)
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

func ObterTodosFaturamentos(c *fiber.Ctx) error {

	pg := ObterPaginacaoReqHttp(c)

	if pg.Ordem == "" {
		pg.Ordem = "referencia desc"
	}

	tb, err := models.ObterTodosFaturamentos(&pg)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "nenhum registro localizado",
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
		"meta":    pg,
	})
}

func AtualizarFaturamento(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id da tabela não informado",
		})
	}

	tb := &models.Faturamento{}

	err := c.BodyParser(&tb)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	tb.ID, _ = uuid.FromString(id)
	err = tb.AtualizarFaturamento()

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

func ApagarFaturamento(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}
	idPesq, _ := uuid.FromString(id)
	err := models.ApagarFaturamento(idPesq)
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

func ObterHistoricoAcumuladoFaturamento(c *fiber.Ctx) error {

	referencia := c.Query("referencia", "")
	if referencia == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}
	deltaMeses, err := strconv.Atoi(c.Query("deltaMeses", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "parametro deltaMeses deve receber apenas numeros",
		})
	}

	totalFaturamento, _, err := models.SomarFaturamentosAnteriores(referencia, deltaMeses)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "nenhum registro localizado",
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "",
		"data": fiber.Map{
			"total": totalFaturamento,
		},
	})
}
