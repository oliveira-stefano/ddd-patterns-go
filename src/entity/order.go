package entity

type Order struct {
	id         string
	customerId string
	items      []OrderItem
	total      float64
}

func (o *Order) NewOrder(id string, customerId string, items []OrderItem) {
	o.id = id
	o.customerId = customerId
	o.items = items
	o.total = o.calculateTotal()
	o.validate()
}

func (o *Order) calculateTotal() float64 {
	total := 0.0
	for _, item := range o.items {
		total += item.GetPrice()
	}
	return total
}

func (o *Order) validate() {
	if o.id == "" {
		panic("Id is empty")
	}
	if o.customerId == "" {
		panic("CustomerId is empty")
	}
	if len(o.items) == 0 {
		panic("Items is empty")
	}
}

func (o *Order) GetTotal() float64 {
	return o.total
}
