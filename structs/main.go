package main

import (
	"fmt"
)

type contactInfo struct {
	email    string
	postCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	oldmate := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:    "partyman3000@email.com",
			postCode: 90210,
		},
	}

	// Assign the memmory pointer useing &amp;
	// oldMatePointer := &oldmate
	oldmate.updateName("Carl")
	oldmate.print()
}

// *person here refers to a pointer of type "person"
func (p *person) updateName(newFirstName string) {
	// *p here refers to the pointer value in memory.
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
