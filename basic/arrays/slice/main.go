package main

import "fmt"

func main() {
	// var numbers []int /* a slice of unspecified size */
	/* numbers == []int{0,0,0,0,0} */
	// numbers = make([]int, 5, 5) /* a slice of length 5 and capacity 5*/

	/* create a slice */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("slice is nil")
	}

	/* print the original slice */
	fmt.Println("numbers ==", numbers)

	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* missing lower bound implies 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* print the sub slice starting from index 0(included) to index 2(excluded) */
	number2 := numbers[:2]
	printSlice(number2)

	/* print the sub slice starting from index 2(included) to index 5(excluded) */
	number3 := numbers[2:5]
	printSlice(number3)

	/* append allows nil slice */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* add one element to slice*/
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* add more than one element at a time*/
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* create a slice numbers1 with double the capacity of earlier slice*/
	numbers4 := make([]int, len(numbers), (cap(numbers))*2)

	/* copy content of numbers to numbers1 */
	copy(numbers4, numbers)
	printSlice(numbers4)
}
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
