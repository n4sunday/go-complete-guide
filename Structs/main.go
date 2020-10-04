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
	sun := person{
		firstName: "Sun",
		lastName:  "Party",
		contact: contactInfo{
			email:   "n4sx@gmail.com",
			zipCode: 80110,
		},
	}

	sunPointer := &sun
	sunPointer.updateName("Nattapon")
	sun.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
