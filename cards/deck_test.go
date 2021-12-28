package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d, ", len(d))
	}
}

func TestFirstAndLastCard(t *testing.T) {
	d := newDeck()

	if d[0] != "♠ Ace ♠" {
		t.Errorf("Expected first card to be [♠ Ace ♠], but got %s", d[0])
	}

	if d[len(d)-1] != "♣ King ♣" {
		t.Errorf("Expected last card to be [♣ King ♣], but got %s", d[len(d)-1])
	}
}

func TestSaveToNewDeckandNewDeckFromFile(t *testing.T) {
	// Need to do all cleanup yourself in Go.
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d, ", len(loadedDeck))
	}

	if loadedDeck[0] != "♠ Ace ♠" {
		t.Errorf("Expected first card to be [♠ Ace ♠], but got %s", loadedDeck[0])
	}

	if loadedDeck[len(d)-1] != "♣ King ♣" {
		t.Errorf("Expected last card to be [♣ King ♣], but got %s", loadedDeck[len(d)-1])
	}

	os.Remove("_decktesting")

}
