package entity

type Order struct {
	id         string
	customerId string
	items      []OrderItem
}

func (o *Order) NewOrder(id string, customerId string, items []OrderItem) {
	o.id = id
	o.customerId = customerId
	o.items = items
}
