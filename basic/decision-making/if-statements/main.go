package main

import "fmt"

func main() {
	ifStatement(10)
	ifElseStatement(100)
	ifElseIfElseStatement(100)
}

func ifStatement(a int) {
	/* local variable definition */
	// var a int = 10
	/* check the boolean condition using if statement */
	if a < 20 {
		/* if condition is true then print the following */
		fmt.Printf("a is less than 20\n")
	}
	fmt.Printf("value of a is : %d\n", a)
}
func ifElseStatement(a int) {
	/* local variable definition */
	// var a int = 100;

	/* check the boolean condition */
	if a < 20 {
		/* if condition is true then print the following */
		fmt.Printf("a is less than 20\n")
	} else {
		/* if condition is false then print the following */
		fmt.Printf("a is not less than 20\n")
	}
	fmt.Printf("value of a is : %d\n", a)
}

func ifElseIfElseStatement(a int) {
	/* local variable definition */
	// var a int = 100

	/* check the boolean condition */
	if a == 10 {
		/* if condition is true then print the following */
		fmt.Printf("Value of a is 10\n")
	} else if a == 20 {
		/* if else if condition is true */
		fmt.Printf("Value of a is 20\n")
	} else if a == 30 {
		/* if else if condition is true  */
		fmt.Printf("Value of a is 30\n")
	} else {
		/* if none of the conditions is true */
		fmt.Printf("None of the values is matching\n")
	}
	fmt.Printf("Exact value of a is: %d\n", a)
}
