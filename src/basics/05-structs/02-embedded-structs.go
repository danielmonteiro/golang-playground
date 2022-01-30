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
	embedded := person{
		firstName: "Embedded",
		lastName:  "Structs",
		contact: contactInfo{
			email:   "embedded-structs@email.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("%+v", embedded)

	// The empty contact will have zero/default values for email and zipcode
	emptyContact := person{firstName: "Daniel", lastName: "Monteiro"}
	fmt.Printf("%+v", emptyContact)
}
