package main

// import the custom package calculator
import (
	"fmt"

	"github.com/sinhnx-dev/golang-examples/tree/main/basic/go-package/custom-pkg/number"
)

func main() {
	i := 5
	fmt.Printf("%d is prime %t\n", i, number.IsPrime(5))
}
