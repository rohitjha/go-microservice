package data

import (
	"time"

	"github.com/google/uuid"
)

// Product defines the structure for an API product
type Product struct {
	ID          uuid.UUID `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	SKU         string    `json:"sku"`
	CreatedOn   time.Time `json:"-"`
	UpdatedOn   time.Time `json:"-"`
}

var productList = []*Product{
	{
		ID:          uuid.New(),
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
	},
	{
		ID:          uuid.New(),
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
	},
}

func ListProducts() []*Product {
	return productList
}

func GetProduct(name string) *Product {
	for _, p := range productList {
		if p.Name == name {
			return p
		}
	}

	return nil
}

func AddProduct(product *Product) {
	product.ID = uuid.New()
	product.CreatedOn = time.Now().UTC()
	product.UpdatedOn = time.Now().UTC()
	productList = append(productList, product)
}

func UpdateProduct(name string, product *Product) {
	for _, p := range productList {
		if p.Name == name {
			p.UpdatedOn = time.Now().UTC()

			if product.Description != "" {
				p.Description = product.Description
			}

			if product.Price != 0.0 {
				p.Price = product.Price
			}

			if product.SKU != "" {
				p.SKU = product.SKU
			}
		}
	}
}

func DeleteProduct(name string) {
	index := 0
	for i, product := range productList {
		if product.Name == name {
			index = i
		}
	}
	productList = append(productList[:index], productList[index+1:]...)
}
