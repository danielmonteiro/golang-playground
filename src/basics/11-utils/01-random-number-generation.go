package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	number := rand.Intn(10)
	fmt.Println("This will always print the same value (because Go uses a default seed for the generator): ", number)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomNumber := r.Intn(10)
	fmt.Println("This will always print random value: ", randomNumber)
}
