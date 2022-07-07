package entity_test

import (
	"testing"

	"github.com/oliveira-stefano/ddd-patterns-go/src/entity"
	"github.com/stretchr/testify/assert"
)

//Order unit test

//Should throw error when id is empty
func TestNewOrder_ShouldThrowErrorWhenIdIsEmpty(t *testing.T) {

	assert.Panics(t, func() {
		order := entity.Order{}
		order.NewOrder("", "123", []entity.OrderItem{})
	}, "Should throw error when id is empty")

}

//Should throw error when customerId is empty
func TestNewOrder_ShouldThrowErrorWhenCustomerIdIsEmpty(t *testing.T) {

	assert.Panics(t, func() {
		order := entity.Order{}
		order.NewOrder("123", "", []entity.OrderItem{})
	}, "Should throw error when customerId is empty")

}

//Should throw error when items is empty
func TestNewOrder_ShouldThrowErrorWhenItemsIsEmpty(t *testing.T) {

	assert.Panics(t, func() {
		order := entity.Order{}
		order.NewOrder("123", "123", []entity.OrderItem{})
	}, "Should throw error when items is empty")

}

//Should calculate total correctly
func TestCalculateTotal_ShouldCalculateTotalCorrectly(t *testing.T) {

	item1 := entity.OrderItem{}
	item1.NewOrderItem("i1", "Item 1", 10.0, "p1", 2)
	item2 := entity.OrderItem{}
	item2.NewOrderItem("i2", "Item 2", 20.0, "p2", 2)
	order := entity.Order{}
	order.NewOrder("123", "123", []entity.OrderItem{item1, item2})
	assert.Equal(t, 60.0, order.GetTotal(), "Should calculate total correctly")

}

//Should throw error if the item quantity is less or equal 0
func TestOrderItem_ShouldThrowErrorIfTheItemQuantityIsLessOrEqual0(t *testing.T) {
	assert.Panics(t, func() {
		item := entity.OrderItem{}
		item.NewOrderItem("i1", "Item 1", 10.0, "p1", 0)
		order := entity.Order{}
		order.NewOrder("123", "123", []entity.OrderItem{item})
	}, "Should throw error if the item quantity is less or equal 0")
}
