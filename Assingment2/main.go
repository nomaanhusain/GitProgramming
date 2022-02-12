package main

import "fmt"

type shape interface {
	area() float64
}
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) area() float64 {
	return 0.5 * t.height * t.base
}
func printArea(sh shape) {
	fmt.Println(sh.area())
}
func main() {
	sqr := square{}
	sqr.sideLength = 4
	printArea(sqr)
	tr := triangle{base: 10, height: 20}
	printArea(tr)
}
