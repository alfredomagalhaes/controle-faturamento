package routes

import (
	"github.com/alfredomagalhaes/controle-faturamento/controllers"
	"github.com/gofiber/fiber/v2"
)

func FechamentoRoute(route fiber.Router) {

	route.Get("/", controllers.ObterTodosFechamentos)
	route.Get("/:id", controllers.ObterFechamento)
	route.Get("/impostos/acumulados", controllers.ObterHistoricoAcumuladoImpostosPagos)

	route.Put("/:id", controllers.AtualizarFechamento)

	route.Post("", controllers.InserirFechamento)
	route.Post("/calcular/:anoMes", controllers.CalcularFechamento)

	route.Delete("/:id", controllers.ApagarFechamento)

}
