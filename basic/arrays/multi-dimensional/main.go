package main

import "fmt"

func main() {
	var threedim [5][10][4]int
	threedim[0][0][0] = 3
	/* an array with 5 rows and 2 columns*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* output each array element's value */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
