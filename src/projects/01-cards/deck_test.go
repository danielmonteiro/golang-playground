package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileReadDeckFromFile(t *testing.T) {
	f := "_decktesting"
	os.Remove(f)

	d := newDeck()

	d.saveToFile(f)

	loadedDeck := readDeckFromFile(f)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(loadedDeck))
	}

	os.Remove(f)
}
