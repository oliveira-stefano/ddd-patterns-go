package entity

type Customer struct {
	id      string
	name    string
	address string `default:""`
	active  bool   `default:"false"`
}

// NewCustomer equivalente ao construtor de Customer
func (c *Customer) NewCustomer(id string, name string) {
	c.id = id
	c.name = name
	c.Validate()
}

func (c *Customer) ChangeName(name string) {
	c.name = name
	c.Validate()
}

func (c *Customer) Activate() {
	if c.address == "" {
		panic("Address is mandatory to activate a customer")
	}
	c.active = true
}

func (c *Customer) Deactivate() {
	c.active = false
}

func (c *Customer) Validate() {
	if c.name == "" {
		panic("Customer name is required")
	} else if c.id == "" {
		panic("Customer id is required")
	}
}
