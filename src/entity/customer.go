package entity

type Customer struct {
	id      string
	name    string
	address Address
	active  bool `default:"false"`
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
	if c.address == (Address{}) {
		panic("Customer address is required")
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

//Create a func SetAddress
func (c *Customer) SetAddress(address Address) {
	c.address = address
}

//Get name
func (c *Customer) GetName() string {
	return c.name
}

//IsActive
func (c *Customer) IsActive() bool {
	return c.active
}
