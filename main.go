package main

import (
	"github.com/rohitjha/go-microservice/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", handlers.APIHandler)
	app.Listen(":3000")
}
