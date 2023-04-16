package main

import "fmt"

/* global variable declaration */
var g int

func main() {
	/* local variable declaration */
	var a, b, c int

	/* actual initialization */
	a = 10
	b = 20
	c = a + b

	fmt.Printf("value of a = %d, b = %d and c = %d\n", a, b, c)
}

/* function to add two integers */
func sum(a, b int) int {
	return a + b
}
