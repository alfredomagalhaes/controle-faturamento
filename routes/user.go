package routes

import (
	"github.com/alfredomagalhaes/controle-faturamento/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router) {

	route.Post("/user", controllers.CreateUser)
}
