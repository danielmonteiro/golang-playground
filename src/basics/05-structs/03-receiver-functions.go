package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	person1 := person{firstName: "Daniel", lastName: "Monteiro"}
	person1.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
