package main

import (
	"os"
	"testing"
)

//if module missing error than run this command: go env -w GO111MODULE=auto
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		//t.Errorf(fmt.Sprint("Expected deck length of 20, but got ", len(d))) //formated string
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades first but got %v", d[0])
	}
	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected Five of Clubs at end but got %v", d[len(d)-1])
	}
}
func TestSaveDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 20 {
		t.Errorf("Loaded deck does not have the appropriate length")
	}
	os.Remove("_decktesting")
}
