package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Gian",
		lastName:  "Gomes",
		contact: contactInfo{
			email:   "any_email@email.com",
			zipCode: 00000,
		},
	}

	jim.updateName("Carlo")
	jim.print()
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
