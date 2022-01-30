package main

import "fmt"

func main() {
	// 1st way to create a map
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "0000FF",
	}
	fmt.Println(colors)

	// 2nd way to create a map
	// var colors1 map[string]string
	colors1 := make(map[string]string)
	colors1["white"] = "#FFFFFF"
	colors1["black"] = "#000000"
	fmt.Println(colors1)

	// Remove a key/value
	delete(colors1, "white")
	fmt.Println(colors1)

	// Iterate over a map
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex color of ", color, "is", hex)
	}
}
