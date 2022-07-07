package entity_test

import (
	"testing"

	"github.com/oliveira-stefano/ddd-patterns-go/src/entity"
	"github.com/stretchr/testify/assert"
)

//Product unit test

//Should throw error when id is empty
func TestProduct_Validate_ShouldThrowErrorWhenIdIsEmpty(t *testing.T) {

	assert.Panics(t, func() {
		product := entity.Product{}
		product.NewProduct("", "Product 1", 10.0)
	}, "Should throw error when id is empty")
}

//Should throw error when name is empty
func TestProduct_Validate_ShouldThrowErrorWhenNameIsEmpty(t *testing.T) {

	assert.Panics(t, func() {
		product := entity.Product{}
		product.NewProduct("1", "", 10.0)
	}, "Should throw error when name is empty")
}

//Should throw error when price is empty
func TestProduct_Validate_ShouldThrowErrorWhenPriceIsEmpty(t *testing.T) {

	assert.Panics(t, func() {
		product := entity.Product{}
		product.NewProduct("1", "Product 1", 0.0)
	}, "Should throw error when price is empty")
}

//Should change name
func TestProduct_ChangeName_ShouldChangeName(t *testing.T) {
	product := entity.Product{}
	product.NewProduct("1", "Product 1", 10.0)
	product.ChangeName("Product 2")
	assert.Equal(t, "Product 2", product.GetName(), "Should change name")
}

//Should change price
func TestProduct_ChangePrice_ShouldChangePrice(t *testing.T) {
	product := entity.Product{}
	product.NewProduct("1", "Product 1", 10.0)
	product.ChangePrice(20.0)
	assert.Equal(t, 20.0, product.GetPrice(), "Should change price")
}
