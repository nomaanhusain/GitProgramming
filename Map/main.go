package main

import "fmt"

func main() {
	// lights := map[string]string{
	// 	"red":    "Stop",
	// 	"yellow": "Ready",
	// 	"green":  "Go",
	// }
	//Amother way
	// var lights map[string] string
	lights := make(map[string]string)
	lights["red"] = "Stop"
	lights["yellow"] = "get Ready"
	lights["green"] = "Go"
	fmt.Println(lights)
	//delete a key value pair delete(map,key)
	delete(lights, "yellow")
	fmt.Println(lights)
	lights["yellow"] = " get Ready"
	fmt.Println(lights["red"])
	printMap(lights)
}
func printMap(m map[string]string) {
	for color, command := range m {
		fmt.Println("The color is ", color, " please ", command)
		fmt.Println()
	}
}
