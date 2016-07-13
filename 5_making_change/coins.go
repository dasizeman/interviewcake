package main

import "fmt"

func main() {
	fmt.Printf("%d ways\n", numberOfWays(4, []int{1, 2, 3}))
}

func numberOfWays(amount int, denominations []int) int {
	numWaysToGetN := make([]int, amount+1)

	// There is one way to get a total of 0 (with no coins)
	numWaysToGetN[0] = 1

	for _, coin := range denominations {
		for greaterAmount := coin; greaterAmount < amount+1; greaterAmount++ {
			// If there is a way to get the current amount with this
			//denomination, it must have at least one of these coins, or we
			// would have counted it already on a previous iteration of the
			// coin loop.  Furthermore, if there is a way that included it any
			// more than once, these will already be recorded in previous calls
			// as well, so we only have to consider the "included once" case
			// each time.  This is the power of the dynamic bottom up approach
			remainder := greaterAmount - coin
			numWaysToGetN[greaterAmount] += numWaysToGetN[remainder]
		}

	}

	return numWaysToGetN[amount]
}
