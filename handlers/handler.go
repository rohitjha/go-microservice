package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func APIHandler(c *fiber.Ctx) error {
	message := "Hello, World!\n"
	log.Print(message)
	c.SendString("Hello, World!")
	return nil
}
