package main

import (
	"fmt"
	"net/http"
	"time"
)

//Unlike the previous 'main.go' code, here we want to check status of all links perpetually
//We will add a gap of few seconds between each request so that we do not oversaturate the servers or look like a DDos attack to the servers
//Will print nothing if no problem, else will print if problem exist
//Make comment executable at line 56 to print STATUS each time
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
	c := make(chan string)
	for _, link := range links {
		go getStatus(link, c)
	}
	//this is an infinite loop
	for {

		//This is a function literal, like lambda function in java.
		//We are using this as we dont want the wait in the first loop
		//If we put this directly in this loop, it will cause the whole main routine to wait for 5 seconds before executing next time
		//While main routine is sleeping, child routines will finish execution and a whole stack of messages will be waiting.
		//This stack will grow very big really soon as the code stops each time and throttle our processing all the time.
		go func(link string) {
			//time.Second is a constant that sleeps the code for 1 sec, we want for 2sec
			time.Sleep(2 * time.Second)
			//This is a blocking statement, will wait for each routine to finish before starting next
			// '<-c' will return link and that will be sent again to getStatus
			getStatus(link, c)
		}(<-c)//Calling this function with required argument
		//this is important, if we pass <-c directly in getStatus, because multiple routines are started here
		//the refrence of <-c might change and hence we need to call by value so that even if refrence changes,
		//the routine started here still work on the same link as when they were started.
		//IMPORTANT RULE - NEVER TRY TO ACCESS A SAME VARIABLE FROM DIFFERENT CHILD ROUTINES, always pass by value or use channels
	}

}
func getStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("***** ", link, "seems to have a problem*****")
		//Put link in channel for 'main' code
		c <- link
		return
	}
	//fmt.Println(link, " OK ")
	//Putting a message in the channel
	c <- link
}
