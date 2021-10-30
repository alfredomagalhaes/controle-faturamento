package controllers

import (
	"fmt"

	"github.com/alfredomagalhaes/controle-faturamento/models"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func InserirFechamento(c *fiber.Ctx) error {

	tb := &models.Fechamento{}

	err := c.BodyParser(&tb)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	err = tb.InserirFechamento()
	if err != nil {
		if err == models.ErrReferenciaFechamentoJaCadastrado {
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

func ObterFechamento(c *fiber.Ctx) error {

	//tb := &models.SN_Tabela{}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}
	idPesq, _ := uuid.FromString(id)
	tb, err := models.ObterFechamento(idPesq)
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

func ObterTodosFechamentos(c *fiber.Ctx) error {

	tb, err := models.ObterTodosFechamentos()
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
	})
}

func AtualizarFechamento(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id da tabela não informado",
		})
	}

	tb := &models.Fechamento{}

	err := c.BodyParser(&tb)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	tb.ID, _ = uuid.FromString(id)
	err = tb.AtualizarFechamento()

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

func ApagarFechamento(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
		})
	}
	idPesq, _ := uuid.FromString(id)
	err := models.ApagarFechamento(idPesq)
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

func CalcularFechamento(c *fiber.Ctx) error {

	referencia := c.Params("anoMes")
	if referencia == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "parametro anoMes não informado",
		})
	}
	totalFat, qtdFat, err := models.SomarFaturamentosAnteriores(referencia, 12)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "malformed request",
			"error":   err,
		})
	}
	//Busca a tabela de faixas de valores do simples nacional
	tabelaSN, _ := models.ObterTabelaVigente(referencia)
	//Busca a tabela de faixas de valores do imposto de renda
	//Busca a tabela de faixas de valores do INSS

	fmt.Println(totalFat)
	fmt.Println(qtdFat)
	fmt.Println(tabelaSN)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "",
	})
}
