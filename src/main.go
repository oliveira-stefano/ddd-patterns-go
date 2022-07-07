package main

import "github.com/oliveira-stefano/ddd-patterns-go/src/entity"

func main() {
	customer := entity.Customer{}
	customer.NewCustomer("123", "John Doe")
	address := entity.Address{}
	address.NewAddress("Rua do Teste", 123, "São Paulo", "012345678")
	customer.SetAddress(address)
	customer.Activate()

	item1 := entity.OrderItem{}
	item1.NewOrderItem("1", "Item 1", 10.0)
	item2 := entity.OrderItem{}
	item2.NewOrderItem("1", "Item 2", 10.0)
	item3 := entity.OrderItem{}
	item3.NewOrderItem("1", "Item 3", 10.0)
	item4 := entity.OrderItem{}
	item4.NewOrderItem("1", "Item 4", 10.0)

	order := entity.Order{}
	order.NewOrder("1", "123", []entity.OrderItem{item1, item2, item3, item4})
}
