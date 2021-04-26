package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProducts(t *testing.T) {
	products := ListProducts()
	assert.Equal(t, 2, len(products))
}

func TestGetProduct(t *testing.T) {
	productNames := []string{"Espresso", "Latte"}

	for _, productName := range productNames {
		product := GetProduct(productName)
		assert.NotNil(t, product)
	}
}

func TestGetProductWhenProductDoesNotExist(t *testing.T) {
	productName := "Pizza"
	product := GetProduct(productName)
	assert.Nil(t, product)
}

func TestAddProduct(t *testing.T) {
	product := &Product{
		Name:        "Pie",
		Description: "Pi",
		Price:       3.1415,
		SKU:         "123456",
	}

	AddProduct(product)

	products := ListProducts()
	assert.Equal(t, 3, len(products))
	foundProduct := GetProduct(product.Name)
	assert.NotNil(t, foundProduct)
	assert.Equal(t, product.Name, foundProduct.Name)
	assert.Equal(t, product.Description, foundProduct.Description)
	assert.Equal(t, product.Price, foundProduct.Price)
	assert.Equal(t, product.SKU, foundProduct.SKU)
}

func TestUpdateProduct(t *testing.T) {
	productName := "Pie"
	productDescription := "Pie"
	product := GetProduct(productName)
	product.Description = productDescription
	UpdateProduct(product.Name, product)

	foundProduct := GetProduct(product.Name)
	assert.NotNil(t, foundProduct)
	assert.Equal(t, productDescription, foundProduct.Description)
}

func TestDeleteProduct(t *testing.T) {
	productName := "Pie"
	DeleteProduct(productName)

	foundProduct := GetProduct(productName)
	assert.Nil(t, foundProduct)
}
