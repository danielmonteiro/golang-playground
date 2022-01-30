package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func main() {
	cards := newDeck()
	fmt.Println("Saving to file: " + cards.toString())
	cards.saveToFile("my_deck")
}

func newDeck() deck {
	return deck{"Ace of Spades", "Two of Hearts", "Three of Diamonds", "Four of Clubs"}
}

func (d deck) toString() string {
	// Converts deck into a slice of string
	str := []string(d)
	return strings.Join(str, ",")
}

func (d deck) saveToFile(filename string) error {
	// Converts the string into a slice of byte (required as second argument of WriteFile function)
	byteSlice := []byte(d.toString())
	return ioutil.WriteFile(filename, byteSlice, 0666)
}
