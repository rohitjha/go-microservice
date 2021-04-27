package main

import (
	"fmt"
	"log"

	"github.com/rohitjha/go-microservice/config"
	"github.com/rohitjha/go-microservice/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()
	app.Get("/products", handlers.ListProducts)
	app.Get("/products/:name", handlers.GetProduct)
	app.Put("/products", handlers.AddProduct)
	app.Post("/products/:name", handlers.UpdateProduct)
	app.Delete("/products/:name", handlers.DeleteProduct)
	return app
}

func main() {
	app := Setup()

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	log.Fatal(
		app.ListenTLS(
			fmt.Sprintf(":%v", config.Port),
			config.CertFileLocation,
			config.KeyFileLocation))
}
