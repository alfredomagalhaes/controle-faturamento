package controllers

import (
	"fmt"
	"strconv"

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

	pg := ObterPaginacaoReqHttp(c)

	if pg.Ordem == "" {
		pg.Ordem = "referencia desc"
	}

	tb, err := models.ObterTodosFechamentos(&pg)
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
	tabelaSN, _ := models.ObterTabelaSNVigente(referencia)
	//Busca a tabela de faixas de valores do imposto de renda
	tabelaIR, _ := models.ObterTabelaIRVigente(referencia)
	//Busca a tabela de faixas de valores do INSS
	tabelaINSS, _ := models.ObterTabelaINSSVigente(referencia)
	//Busca o faturamento do mês
	fatMes, _ := models.ObterFaturamentoMes(referencia)

	//Se possuir menos que 12 meses de faturamento
	//o total faturado será a média dos meses * 12
	if qtdFat < 12 {
		totalFat = (totalFat / qtdFat) * 12
	}
	//verfica qual a faixa de valor que o faturamento se encontra
	idxSN := -1
	for i, fx := range tabelaSN.Faixas {
		if totalFat >= fx.ValorInicial && totalFat <= fx.ValorFinal {
			idxSN = i
			break
		}
	}
	if idxSN < 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "",
		})
	}
	aliqSN := (((totalFat * tabelaSN.Faixas[idxSN].Aliquota) - tabelaSN.Faixas[idxSN].ParcelaDeducao) / totalFat)
	targetProLabore := fatMes.ValorFaturado * (tabelaSN.PercentualFolha / 100)

	//verifica a faixa de valor do INSS
	idxINSS := -1
	for i, fx := range tabelaINSS.Faixas {
		if targetProLabore >= fx.ValorInicial && targetProLabore <= fx.ValorFinal {
			idxINSS = i
			break
		}
	}
	valorInss := 0.00
	if idxINSS == -1 {
		valorInss = tabelaINSS.ValorLimite
	} else {
		valorInss = (targetProLabore * (tabelaINSS.Faixas[idxINSS].Aliquota / 100))
		if valorInss > tabelaINSS.ValorLimite {
			valorInss = tabelaINSS.ValorLimite
		}
	}

	//Verifica a faixa de valores do imposto de renda
	idxIRRF := -1
	for i, fx := range tabelaIR.Faixas {
		if targetProLabore >= fx.ValorInicial && targetProLabore <= fx.ValorFinal {
			idxIRRF = i
			break
		}
	}
	//calcula o valor do IR a ser pago
	valorIR := ((targetProLabore * (tabelaIR.Faixas[idxIRRF].Aliquota / 100)) - tabelaIR.Faixas[idxIRRF].ParcelaDeducao)
	valorIR -= (2 * tabelaIR.DeducaoDependente) //TODO - adicionar tabela de depenentes do IR
	if valorIR < 0 {
		valorIR = 0
	}
	//Calcula o valor da DAS
	valorDAS := fatMes.ValorFaturado * (aliqSN / 100)

	fechaMes := models.Fechamento{}
	fechaMes.Referencia = referencia
	fechaMes.ValorDAS, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", valorDAS), 64)
	fechaMes.AliquotaDAS = aliqSN
	fechaMes.ValorINSS, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", valorInss), 64)
	fechaMes.AliquotaINSS = tabelaINSS.Faixas[idxINSS].Aliquota
	fechaMes.ValorIRRF, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", valorIR), 64)
	fechaMes.AliquotaIRRF = tabelaIR.Faixas[idxIRRF].Aliquota
	fechaMes.Tipo = "P"

	buscaF, _ := models.ObterFechamentoPorMesETipo(fechaMes.Referencia, fechaMes.Tipo)
	if buscaF.Referencia != "" {
		fechaMes.ID = buscaF.ID
		err = fechaMes.AtualizarFechamento()
	} else {

		err = fechaMes.InserirFechamento()
	}
	if err != nil {
		return c.Status(fiber.StatusPreconditionFailed).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "",
		"data":    fechaMes,
	})
}

func ObterHistoricoAcumuladoImpostosPagos(c *fiber.Ctx) error {

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

	totalImpostos, err := models.SomarImpostosPagos(referencia, deltaMeses)
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
			"total": totalImpostos,
		},
	})

}
