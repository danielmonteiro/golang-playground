package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	fmt.Println("--- HAND ---")
	hand, remainingDeck := deal(cards, 5)
	hand.print()

	fmt.Println("--- REMAINING DECK ---")
	remainingDeck.print()
}
