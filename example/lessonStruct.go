package main

import "log"

type MyConstraint interface {
	int | float64
}

type customer struct {
	name    string
	address string
	bal     float64
}

type contact struct {
	firstName   string
	lastName    string
	phoneNumber string
}

type business struct {
	name    string
	address string
	contact
}
type rectangle struct {
	length, height float64
}

func main() {
	var info customer
	info.name = "Tom Smith"
	info.address = "5 main St"
	info.bal = 234

	log.Println("5 + 4 = ", getSumGen(5, 8))
	log.Println("5 + 4 = ", getSumGen(10.6, 2.56))
	getCustInfo(info)
	newCustAdd(&info, "Verdi 6")
	log.Println("Address :", info.address)

	sCustomer := customer{name: "Jelly Smith", address: "123 Main", bal: 130.4}
	log.Println("Second customer Name :", sCustomer.name)

	rect := rectangle{10.0, 15.0}
	log.Println("React area :", rect.Area())

	con := contact{
		"James", "Wang", "555-444-333",
	}
	bus := business{
		"ECO", "234 North St", con,
	}
	bus.info()
}

func (b business) info() {
	log.Println("Contact at ", b.name, b.contact.firstName, b.phoneNumber)
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

func getCustInfo(c customer) {
	log.Println("Name :", c.name)

}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}
