package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 8)
	fmt.Println("RemainingDeck = ")
	remainingDeck.printDeck()
	hand.saveToFile("my_hand")
	fmt.Println("From file")
	fmt.Println(newDeckFromFile("my_hand"))
	fmt.Println("Shuffle hand")
	hand.shuffle()
	hand.printDeck()
}
