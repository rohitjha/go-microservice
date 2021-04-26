package main

import (
	"log"

	"github.com/rohitjha/go-microservice/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()
	app.Get("/", handlers.APIHandler)
	return app
}

func main() {
	app := Setup()
	log.Fatal(app.Listen(":3000"))
}
