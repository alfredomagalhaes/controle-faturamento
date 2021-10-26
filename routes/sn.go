package routes

import (
	"github.com/alfredomagalhaes/controle-faturamento/controllers"
	"github.com/gofiber/fiber/v2"
)

func SNRoute(route fiber.Router) {

	route.Post("", controllers.CriarTabelaSN)
	route.Get("/:id", controllers.ObterTabelaSN)

}
