package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://fast.com",
		"http://twitter.com",
		"http://leetcode.com",
		"http://amazon.in",
	}
	//make(chan <type>) is used to create a channel for inter-routine commmunication.
	//We will use string to communicate hence string is type
	c := make(chan string)
	for _, link := range links {
		//go keyword is used to create child routines inside the main routine, this introduces concurrency
		//Still main routine has maximum priority, hence program quits even before all of the child routines finish executing
		//We need to use channels to comunicate between different routines
		go getStatus(link, c)
	}
	for i := 0; i < len(links); i++ {
		//This makes the main routine listen for messages on channel c
		//This will cause it to wait until a message is recieved and only then will it continue execution
		//To listen for a message on a channel we use this syntax "message <- c"
		fmt.Println(<-c) //This is a blocking statement
	}

}
func getStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("***** ", link, "seems to have a problem*****")
		return
	}
	fmt.Println(link, " OK ")
	//Putting a message in the channel
	c <- "I am done!"
}
