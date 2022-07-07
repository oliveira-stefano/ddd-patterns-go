package entity

type OrderItem struct {
	id        string
	productId string
	name      string
	price     float64
	quantity  int
}

func (o *OrderItem) NewOrderItem(id string, name string, price float64, productId string, quantity int) {
	o.id = id
	o.name = name
	o.price = price
	o.productId = productId
	o.quantity = quantity
}

func (o *OrderItem) GetPrice() float64 {
	return o.price * float64(o.quantity)
}
