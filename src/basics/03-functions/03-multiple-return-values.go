package main

import "fmt"

func main() {
	fruit1, fruit2 := getFruits()
	fmt.Println("Fruit 1: " + fruit1 + ", Fruit 2: " + fruit2)
}

func getFruits() (string, string) {
	return "apple", "banana"
}
