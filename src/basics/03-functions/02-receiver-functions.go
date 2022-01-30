package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func main() {
	cards := deck{"Ace of Spades", newCard(), "Five of Diamonds"}
	cards.print()
}

func newCard() string {
	return "Two of Diamonds"
}
