package routes

import (
	"github.com/alfredomagalhaes/controle-faturamento/controllers"
	"github.com/gofiber/fiber/v2"
)

func SNRoute(route fiber.Router) {

	route.Post("", controllers.CriarTabelaSN)
	route.Get("/", controllers.ObterTodasTabelasSN)
	route.Get("/:id", controllers.ObterTabelaSN)
	//route.Get("/vigente", controllers.ObterTabelaSNVigente)
	route.Put("/:id", controllers.AtualizarTabelaSN)
	route.Get("/:id/faixas", controllers.ObterFaixasPorTabelaSN)
	route.Delete("/faixas/:id", controllers.ApagarFaixaSN)

}
