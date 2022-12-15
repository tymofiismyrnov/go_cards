package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 52
	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("The first card is expected to be Ace of Spades, but go %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("The last card is expected to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected %v cards in deck, got %v", len(deck), len(loadedDeck))
	}

	os.Remove("_decktesting")
}
