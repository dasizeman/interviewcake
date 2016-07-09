/*
INDEPENDENT SOLUTION
From now on, I am going to post seperately the solutions I come up with
initially (independently), and then maybe an improved version once I have
looked at the InterviewCake solution for archiving's sake.

What I learned:
Although this solution is O(n) time and O(1) additional space, I could have
been even more greedy in how I keep track of if the current item contributes to a high product or not, by storing highestProductOf2, highest, and lowest.  I
also relied on my previous solution's algorithm (but code re-use is good right?)

At the very least this algorithm can very easily be expanded to "highest
product of x"
*/
package main

import (
	"fmt"
	"log"
)

func main() {

	values := []int{-10, -10, 1, 3, 2}
	fmt.Printf("You entered: %v\n", values)

	fmt.Printf("Best product: %d\n", highestProductOf3(values))

}

func highestProductOf3(values []int) int {
	currentBestMultiplicands := make([]int, 3)
	currentBestProduct := 1

	if len(values) < 3 {
		log.Fatalf("highestProductOf3() must be passed a slice of at least 3 values\n")
	}
	copy(currentBestMultiplicands, values[:3])

	fmt.Printf("%v\n", currentBestMultiplicands)

	// Calculate an initial "best" value
	for _, val := range currentBestMultiplicands {
		currentBestProduct *= val
	}

	for i := 3; i < len(values); i++ {

		// Determine if our product would be higher if we replaced any of our
		// current best multiplicands
		var potentialProduct, potentialReplaceIdx int = -1, -1

		// Use one of our old solutions to get the products without
		// each respective item
		partialProducts := getProductsOfAllIntsExceptAtIndex(currentBestMultiplicands)

		for j, val := range partialProducts {
			potentialProduct = values[i] * val
			if potentialProduct > currentBestProduct {
				// Keep track of the best product we've seen so far
				// We can only replace one of our "best values" with the
				// current value, so we replace the one that results in the
				// highest product
				currentBestProduct = potentialProduct
				potentialReplaceIdx = j
			}
		}

		// If we have determined that making a replacement gives a better
		// product, replace the item
		if potentialReplaceIdx >= 0 {
			currentBestMultiplicands[potentialReplaceIdx] = values[i]
		}

	}

	return currentBestProduct

}

func getProductsOfAllIntsExceptAtIndex(a []int) (res []int) {
	// First we run through the input and create a data structure that holds
	// the product of all values up to that index, one for each direction
	products := make([]int, len(a))
	products[0] = 1

	currentProduct := 1

	// In the following loops we compute the result values by greedily finding
	// the product of all values up to and not including the value at the
	// current index in both directions, and multiplying these values together
	for i := 0; i < len(a); i++ {
		products[i] = currentProduct
		currentProduct *= a[i]
	}

	res = make([]int, len(a))
	currentProduct = 1
	for i := len(a) - 1; i >= 0; i-- {
		res[i] = products[i] * currentProduct
		currentProduct *= a[i]
	}

	return
}
