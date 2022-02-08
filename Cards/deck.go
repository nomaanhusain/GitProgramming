package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Custom data types called 'deck' which is a slice of string
type deck []string //deck is new type that is equvalent to a slice of string

//func to create a new deck and return it
func newDeck() deck {
	cards := deck{}

	//we are using a smart way to build out the cards as we don't want to write names of 52 cards
	//we can use nested loops to make combination of all suits and values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}
	// '_' symbol is used to tell go that we dont care about this variable
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//func to deal out a deck of specified size, it will return 2 deck (slices), one of specified size, and one of left overs
//function with 2 parameters and 2 return values
func deal(d deck, handSize int) (deck, deck) {
	//split a slice using ':', d[inclusiveValue:nonInclusiveValue]
	//ADD LENGTH CHECK AND PROPER ERROR MESSAGE IF HANDSIZE IS BIGGER THAN DECK
	return d[:handSize], d[handSize:]
}

//func to save a deck to file
func (d deck) saveToFile(filename string) error {
	//0666 is a permission that allows read and write
	//WriteFile function returns error, so we return it
	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
}

//func to create new deck from filename
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		//if there is error, print it and exit program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}

//func to shuffle a deck
func (d deck) shuffle() {
	//create out own rand object using current time as seed as that will always be changing
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i, _ := range d {
		//generate random and swap
		//newPos := rand.Intn(len(d) - 1) //this is not truly random and comes from a static seed value
		newPos := r.Intn(len(d) - 1)
		t := d[i]
		d[i] = d[newPos]
		d[newPos] = t

		// or d[newPos], d[i] = d[i], d[newPos]
	}
}

//func (reciever) funcname(){}
//func to print all cards
func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//deck to string
//we will join all strings to a single string of comma seperated values
func (d deck) deckToString() string {
	//function from strings lib
	return strings.Join([]string(d), ",")
}
