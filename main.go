package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func api_handler(c *fiber.Ctx) error {
	message := "Hello, World!\n"
	log.Print(message)
	return c.SendString("Hello, World!")
}

func main() {
	app := fiber.New()
	app.Get("/", api_handler)
	app.Listen(":3000")
}
