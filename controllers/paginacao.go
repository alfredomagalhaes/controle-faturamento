package controllers

import (
	"strconv"

	"github.com/alfredomagalhaes/controle-faturamento/models"
	"github.com/gofiber/fiber/v2"
)

func ObterPaginacaoReqHttp(c *fiber.Ctx) models.Paginacao {

	pg := models.Paginacao{}

	pagina, _ := strconv.Atoi(c.Query("pagina"))
	if pagina == 0 {
		pagina = 1
	}
	pg.Pagina = pagina

	porPagina, _ := strconv.Atoi(c.Query("por_pagina"))
	if porPagina == 0 {
		porPagina = 10 //TODO inserir variavel default
	}
	pg.Limite = porPagina

	pg.Ordem = c.Query("ordem")

	return pg
}
