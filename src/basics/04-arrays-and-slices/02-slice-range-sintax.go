package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "grape", "orange"}
	fmt.Println(fruits[0:2])
	fmt.Println(fruits[:2])
	fmt.Println(fruits[2:])
}
