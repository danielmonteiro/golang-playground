package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func main() {
	cards := readDeckFromFile("my_deck")
	cards.print()
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func readDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
