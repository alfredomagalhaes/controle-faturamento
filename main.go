package main

import (
	"log"
	"os"

	"github.com/alfredomagalhaes/controle-faturamento/config"
	"github.com/alfredomagalhaes/controle-faturamento/models"
	"github.com/alfredomagalhaes/controle-faturamento/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// config db
	config.ConnectDB()

	//Cria as tabelas a serem usadas no banco de dados
	models.InicializarTabelas()

	// setup routes
	setupRoutes(app)

	// Listen on server 8000 and catch error if any
	err = app.Listen(":8000")

	// handle error
	if err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉 🇧🇷",
		})
	})

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// Login endpoint
	routes.LoginRoute(api.Group("/login"))

	routes.UserRoute(api)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("token_password")),
	}))

	//Every route declared below this point, will be using JWT authentication
	routes.SNRoute(api.Group("/tabelaSN"))
	routes.IRRoute(api.Group("/tabelaIR"))
	routes.INSSRoute(api.Group("/tabelaINSS"))
	routes.FaturamentoRoute(api.Group("/faturamento"))
	routes.FechamentoRoute(api.Group("/fechamento"))

}
