package main

import "fmt"

func main() {
	var students map[string]string // Can be used to create empty map
	fmt.Println(students)
	clothes := make(map[string]string) // Can also be used to create empty map
	fmt.Println(clothes)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)

	}
}
