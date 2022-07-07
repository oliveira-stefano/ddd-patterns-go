package entity

type OrderItem struct {
	id    string
	name  string
	price float64
}

func (o *OrderItem) NewOrderItem(id string, name string, price float64) {
	o.id = id
	o.name = name
	o.price = price
}
