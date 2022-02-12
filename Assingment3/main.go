package main

import (
	"fmt"
	"io"
	"os"
)

//run this by 'go run main.go nomaan.txt'
func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)
	ourFile, _ := os.Open(fileName)
	io.Copy(os.Stdout, ourFile)
}
