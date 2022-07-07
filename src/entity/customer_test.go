package entity_test

//Customer unit test
import (
	"testing"

	"github.com/oliveira-stefano/ddd-patterns-go/src/entity"
	"github.com/stretchr/testify/assert"
)

//Should panic when id is empty
func TestCustomer_NewCustomer_ShouldPanicWhenIdIsEmpty(t *testing.T) {
	assert.Panics(t, func() {
		customer := entity.Customer{}
		customer.NewCustomer("", "John Doe")
	}, "Should panic when id is empty")
}

//Should panic when name is empty
func TestCustomer_NewCustomer_ShouldPanicWhenNameIsEmpty(t *testing.T) {
	assert.Panics(t, func() {
		customer := entity.Customer{}
		customer.NewCustomer("123", "")
	}, "Should panic when name is empty")
}

//Should change name
func TestCustomer_ChangeName_ShouldChangeName(t *testing.T) {
	customer := entity.Customer{}
	customer.NewCustomer("123", "John")
	customer.ChangeName("John Doe")
	assert.Equal(t, "John Doe", customer.GetName(), "Should change name")
}

//Should activate customer
func TestCustomer_Activate_ShouldActivateCustomer(t *testing.T) {
	customer := entity.Customer{}
	customer.NewCustomer("123", "John")
	address := entity.Address{}
	address.NewAddress("Rua do Teste", 123, "São Paulo", "012345678")
	customer.SetAddress(address)
	customer.Activate()
	assert.True(t, customer.IsActive(), "Should activate customer")
}

//Should deactivate customer
func TestCustomer_Deactivate_ShouldDeactivateCustomer(t *testing.T) {
	customer := entity.Customer{}
	customer.NewCustomer("123", "John")
	customer.Deactivate()
	assert.False(t, customer.IsActive(), "Should deactivate customer")
}

//should throw error when address is empty
func TestCustomer_Activate_ShouldThrowErrorWhenAddressIsEmpty(t *testing.T) {
	assert.Panics(t, func() {
		customer := entity.Customer{}
		customer.NewCustomer("123", "John")
		customer.Activate()
	}, "Should throw error when address is empty")
}
