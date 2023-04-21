package main

// import the custom package calculator
import (
	"fmt"

	"github.com/sinhnx-dev/golang-examples/basic/go-package/custome-pkg/number"
)

func main() {
	i := 5
	fmt.Printf("%d is prime %t\n", i, number.IsPrime(5))
}
