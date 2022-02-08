package main

import "fmt"

func main1() {
	var cardName string = "Ace of Spades" //method 1 for delare and assign variables
	cardName1 := "Ace of Spades"          //method 2 for delare and assign variable
	cardName1 = "Five of Hearts"          // reassign variable
	fmt.Println(cardName, cardName1)
	cardName = newCard()
	fmt.Println(cardName)

	//Making a slice, Slice is like an array but if mutable length
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "The trump card") //append does not modify, but returns new slice,
	//so you have to put previous slice inside it
	fmt.Println(cards)
	//iteration over slice
	for index, crd := range cards {
		fmt.Println(index, crd)
	}
}

//new function
//syntax = func 'name'() 'return type'{}
func newCard() string {
	return "Three of Diamonds"
}
