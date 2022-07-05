package entity

type Customer struct {
	id      string
	name    string
	address string
	active  bool
}

// NewCustomer equivalente ao construtor de Customer
func (c *Customer) NewCustomer(id string, name string, address string, active bool) {
	c.id = id
	c.name = name
	c.address = address
	c.active = active
}

func (c *Customer) ChangeName(name string) {
	c.name = name
}

func (c *Customer) Activate() {
	c.active = true
}

func (c *Customer) Deactivate() {
	c.active = false
}
