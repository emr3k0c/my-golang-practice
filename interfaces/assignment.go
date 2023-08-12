package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	tri := triangle{
		height: 15,
		base:   17,
	}
	squ := square{
		sideLength: 20,
	}
	printArea(tri)
	printArea(squ)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
