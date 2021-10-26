package routes

import (
	"github.com/alfredomagalhaes/controle-faturamento/controllers"
	"github.com/gofiber/fiber/v2"
)

func LoginRoute(route fiber.Router) {

	route.Post("", controllers.Login)
}
