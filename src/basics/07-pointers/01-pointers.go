package main

import "fmt"

type person struct {
	name  string
	email string
}

func main() {
	/*
	 * This will not update values as the function is updating a copy
	 */
	p := person{
		name:  "Daniel",
		email: "daniel@email.com",
	}
	fmt.Printf("%+v\n", p)
	p.updateName("New Name")
	fmt.Printf("%+v\n", p)

	fmt.Println("=======================================")

	/*
	 * We can obtain the pointer of a variable using &, and then with
	 * the pointer we can update the "real" variable inside a function
	 */
	personPointer := &p
	fmt.Println(personPointer)
	personPointer.updateNamePointer("New Name with Pointer")
	fmt.Printf("%+v\n", p)

	fmt.Println("=======================================")

	/*
	 * Go provides a shortcut that it automatically converts to a pointer
	 */
	p.updateNamePointer("Shortcut")
	fmt.Printf("%+v\n", p)

	fmt.Println("=======================================")

	/*
	 * Slices can be updated inside a function, because a slice is a struct that
	 * points to an array under the hood
	 */
	s := []string{"Hello", "World"}
	updateSlice(s)
	fmt.Println(s)
}

func (p person) updateName(newName string) {
	p.name = newName
}

func (p *person) updateNamePointer(newName string) {
	(*p).name = newName
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
