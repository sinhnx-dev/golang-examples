package main

import "fmt"

func main() {
	// var a int = 10
	// fmt.Printf("Address of a variable: %x\n", &a)

	// var ip *int     /* pointer to an integer */
	// var fp *float32 /* pointer to a float */

	var a int = 20 /* actual variable declaration */
	var ip *int    /* pointer variable declaration */

	ip = &a /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	var ptr *int
	fmt.Printf("The value of ptr is : %x\n", ptr)
}
