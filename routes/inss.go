package routes

import (
	"github.com/alfredomagalhaes/controle-faturamento/controllers"
	"github.com/gofiber/fiber/v2"
)

func INSSRoute(route fiber.Router) {

	route.Post("", controllers.CriarTabelaINSS)
	route.Get("/:id", controllers.ObterTabelaINSS)
	route.Put("/:id", controllers.AtualizarTabelaINSS)
	route.Get("/:id/faixas", controllers.ObterFaixasPorTabelaINSS)
	route.Delete("/faixas/:id", controllers.ApagarFaixaINSS)

}
