package main

import (
	"log"

	"github.com/rohitjha/go-microservice/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()
	app.Get("/products", handlers.GetProducts)
	app.Put("/products", handlers.AddProduct)
	app.Post("/products/:name", handlers.UpdateProduct)
	app.Delete("/products/:name", handlers.DeleteProduct)
	return app
}

func main() {
	app := Setup()
	log.Fatal(app.ListenTLS(":3001", "certs/localhost.crt", "certs/localhost.key"))
}
