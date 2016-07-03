package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", rand7())
	}
}

func rand7() int {

	outcome := 21
	// Generate an outcome of the "roll" of two rand5() calls, throw out values
	// above 21 (we are using zero based indexing for outcomes)

	for outcome > 20 {
		rollone := rand5()
		rolltwo := rand5()

		outcome = (rollone-1)*5 + (rolltwo - 1)
	}

	return outcome/3 + 1
}

// Generate a random number between 1 and 5
func rand5() int {
	return rand.Intn(5) + 1
}
