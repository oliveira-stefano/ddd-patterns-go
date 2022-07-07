package entity

import "strconv"

type Address struct {
	street string
	number int `default:"0"`
	city   string
	zip    string
}

func (a *Address) NewAddress(street string, number int, city string, zip string) {
	a.street = street
	a.number = number
	a.city = city
	a.zip = zip
	a.Validate()
}

func (a *Address) Validate() {
	if a.street == "" {
		panic("Street is required")
	} else if a.number == 0 {
		panic("Number is required")
	} else if a.city == "" {
		panic("City is required")
	} else if a.zip == "" {
		panic("Zip is required")
	}
}

func (a *Address) ToString() string {
	return a.street + " " + strconv.Itoa(a.number) + " " + a.city + " " + a.zip
}
