package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rohitjha/go-microservice/data"
)

func ListProducts(c *fiber.Ctx) error {
	log.Print("Getting all products")
	products := data.ListProducts()
	c.Status(fiber.StatusOK)
	c.JSON(products)
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	name := c.Params("name")
	log.Printf("Getting product %v", name)
	product := data.GetProduct(name)

	if product == nil {
		log.Printf("Did not find any product with name %v", name)
		c.Status(fiber.StatusNotFound)
		return nil
	}

	c.Status(fiber.StatusOK)
	c.JSON(product)
	return nil
}

func AddProduct(c *fiber.Ctx) error {
	product := new(data.Product)
	c.BodyParser(product)
	log.Printf("Adding product %v", product.Name)
	data.AddProduct(product)
	c.Status(fiber.StatusCreated)
	return nil
}

func UpdateProduct(c *fiber.Ctx) error {
	name := c.Params("name")
	log.Printf("Updating product %v", name)
	product := new(data.Product)
	c.BodyParser(product)
	data.UpdateProduct(name, product)
	c.Status(fiber.StatusOK)
	return nil
}

func DeleteProduct(c *fiber.Ctx) error {
	name := c.Params("name")
	log.Printf("Deleting product %v", name)
	data.DeleteProduct(name)
	c.Status(fiber.StatusNoContent)
	return nil
}
