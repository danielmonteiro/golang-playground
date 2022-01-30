package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// 1st way of declaring structs (arguments follows same order as field declaration)
	person1 := person{"Daniel", "Monteiro"}
	fmt.Printf("%+v", person1)

	// 2nd way of declaring structs (assign values to fields individually)
	person2 := person{firstName: "Daniel", lastName: "Monteiro"}
	fmt.Printf("%+v", person2)

	// 3rd way of declaring structs (empty struct with "zero/default values" for fields)
	var person3 person
	person3.firstName = "Daniel"
	person3.lastName = "Monteiro"
	fmt.Printf("%+v", person3)
}
