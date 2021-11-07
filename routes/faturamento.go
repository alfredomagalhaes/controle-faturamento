package routes

import (
	"github.com/alfredomagalhaes/controle-faturamento/controllers"
	"github.com/gofiber/fiber/v2"
)

func FaturamentoRoute(route fiber.Router) {

	route.Post("", controllers.InserirFaturamento)
	route.Get("/", controllers.ObterTodosFaturamentos)
	route.Get("/:id", controllers.ObterFaturamento)
	route.Put("/:id", controllers.AtualizarFaturamento)
	route.Delete("/:id", controllers.ApagarFaturamento)
	route.Get("/historico/acumulado", controllers.ObterHistoricoAcumuladoFaturamento)

}
