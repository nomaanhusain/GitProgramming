package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	//making a response
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string][]string)
		//calling function
		downStatus := hitStatus()
		if len(downStatus) == 0 {
			resp["webstatus"] = []string{"All Active"}
			resp["down"] = []string{""}
		} else {
			resp["webstatus"] = []string{"Problem"}
			//resp["down"] = strings.Join(downStatus, " ")
			resp["down"] = downStatus
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	}
	//We have used and external library for handeling cors https://github.com/rs/cors
	//As we want cross origin enabled so that this API can be accessed by other applications line 36,37,38 is for that
	mux := http.NewServeMux()
	mux.HandleFunc("/", h1)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
func hitStatus() []string {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://fast.com",
		"http://twitter.com",
		"http://leetcode.com",
		"http://amazon.in",
		"http://amazon.com",
		"http://netflix.com",
		"http://hulu.com",
		"http://twitch.tv",
		"http://youtube.com",
	}
	c := make(chan bool)
	messg := make(chan string)
	for _, link := range links {
		go getStatus(link, c, messg)
	}
	problemLinks := []string{}
	for i := 0; i < len(links); i++ {
		x := <-c
		if !(x) {
			problemLinks = append(problemLinks, <-messg)
			continue
		}
		fmt.Println(<-messg)
	}
	return problemLinks
}
func getStatus(link string, c chan bool, messg chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("***** ", link, "seems to have a problem*****")
		c <- false
		messg <- link
		return
	}
	//fmt.Println(link, " OK ")
	//Putting a message in the channel
	c <- true
	messg <- "OK"
}
