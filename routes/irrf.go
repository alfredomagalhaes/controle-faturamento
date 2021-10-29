package routes

import (
	"github.com/alfredomagalhaes/controle-faturamento/controllers"
	"github.com/gofiber/fiber/v2"
)

func IRRoute(route fiber.Router) {

	route.Post("", controllers.CriarTabelaIR)
	route.Get("/:id", controllers.ObterTabelaIRRF)
	route.Put("/:id", controllers.AtualizarTabelaIR)
	route.Get("/:id/faixas", controllers.ObterFaixasPorTabelaIR)
	route.Delete("/faixas/:id", controllers.ApagarFaixaIR)

}
