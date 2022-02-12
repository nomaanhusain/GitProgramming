package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://reqres.in/api/users/2")
	if err != nil {
		fmt.Println("error : ", err)
		os.Exit(1)
	}
	//One way to do this
	// //(type, capacity)
	// bs := make([]byte, 99999)
	// //fill bs with Body of resp
	// //Read is Reader Interface, see io reader documentation https://pkg.go.dev/io@go1.17.7#Reader
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//Another way is using io.Copy
	//Read documentation https://pkg.go.dev/io@go1.17.7#Copy
	//1st argument has to implement writer interface https://pkg.go.dev/io@go1.17.7#Writer
	//2nd argument has to implement reader interface
	//pull data from 2nd argument and put it in 1st argument
	io.Copy(os.Stdout, resp.Body)
}
