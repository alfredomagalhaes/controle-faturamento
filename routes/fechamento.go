package routes

import (
	"github.com/alfredomagalhaes/controle-faturamento/controllers"
	"github.com/gofiber/fiber/v2"
)

func FechamentoRoute(route fiber.Router) {

	route.Post("", controllers.InserirFechamento)
	route.Get("/", controllers.ObterTodosFechamentos)
	route.Get("/:id", controllers.ObterFechamento)
	route.Put("/:id", controllers.AtualizarFechamento)
	route.Delete("/:id", controllers.ApagarFechamento)

}
