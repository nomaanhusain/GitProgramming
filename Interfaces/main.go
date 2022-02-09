package main

import "fmt"

//this is an interface
type bot interface {
	//any function with name 'getGreeting' and return type string are now like a child of this interface
	//so type bot can accept values of type englishBot and spanishBot acc. to this implementation
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

//bot can accept values of type englishBot and spanishBot
func printGreeting(b bot) {
	fmt.Print(b.getGreeting())
	fmt.Println()
}

//we can remove variable name from reciever because we are not using it inside function, it is alright
func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
