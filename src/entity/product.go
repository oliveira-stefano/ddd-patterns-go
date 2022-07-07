package entity

type Product struct {
	id    string
	name  string
	price float64
}

//NewProduct as constructor
func (p *Product) NewProduct(id string, name string, price float64) {
	p.id = id
	p.name = name
	p.price = price
	p.validate()
}

func (p *Product) validate() {
	if p.id == "" {
		panic("Product id is empty")
	}
	if p.name == "" {
		panic("Product name is empty")
	}
	if p.price <= 0 {
		panic("Product price is invalid")
	}
}

func (p *Product) ChangePrice(price float64) {
	p.price = price
}

func (p *Product) ChangeName(name string) {
	p.name = name
	p.validate()
}

func (p *Product) GetName() string {
	return p.name
}

func (p *Product) GetPrice() float64 {
	return p.price
}
