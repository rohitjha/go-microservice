package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func APIHandler(c *fiber.Ctx) error {
	status := fiber.StatusOK
	message := "Hello, World!"
	log.Print(message)
	c.Status(status)
	c.JSON(fiber.Map{"status": status, "message": message, "data": nil})
	return nil
}
