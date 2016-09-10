package main

import (
	"fmt"
	"math"
)

// CakeStat represents the weight and value of a cake
type CakeStat struct {
	weight, value int
}

func main() {
	cakeTypes := []CakeStat{
		{7, 160},
		{3, 90},
		{2, 15},
		{0, 0},
		{15, 0}}

	capacity := 20

	fmt.Printf("Max value: %d\n", maxDuffelBagValue(cakeTypes, capacity))
}

func maxDuffelBagValue(cakeTypes []CakeStat, bagCapacity int) int {

	// How much weight we have yet to fill
	remainingWeight := bagCapacity

	// The maximum value of our bag
	maxBagValue := 0

	// A convenient way to keep track of which cake types have
	// been "used" so we don't have to modify the list
	used := make(map[int]bool)

	bagFull := false

	for !bagFull {
		// The index of the cake type that we can currently gain the most value
		// from
		bestValueIdx := 0

		// How much value we gain by filling our bag with the maximum amount of
		// the aformentioned cake type
		maxValueAdded := 0

		// How many cakes we have just added
		maxNumCakesAdded := 0

		// Find the most value we can get from an unused type of cake
		for idx, cake := range cakeTypes {
			// If we have already used this cake type, skip it
			_, isUsed := used[idx]
			if isUsed {
				continue
			}

			// If this cake has zero value, it isn't useful
			if cake.value == 0 {
				used[idx] = true
				continue
			}

			// If we have a cake with nonzero value but zero weight, our maximum
			// bag value is infinite
			if cake.weight == 0 {
				return math.MaxInt32
			}

			// If the weight of this cake is higher than our remaining bag
			// weight, we can't use it
			if cake.weight > remainingWeight {
				used[idx] = true
				continue
			}

			// What is the max value we can add with this cake type?
			numCakesPossible := (remainingWeight / cake.weight)
			possibleValueAdded := numCakesPossible * cake.value
			if possibleValueAdded > maxValueAdded {
				bestValueIdx = idx
				maxValueAdded = possibleValueAdded
				maxNumCakesAdded = numCakesPossible
			}
		}

		// If we can't further increase the bag value, we're done
		if maxValueAdded == 0 {
			bagFull = true
			break
		}

		// "Add the cakes to our bag"
		remainingWeight -= cakeTypes[bestValueIdx].weight * maxNumCakesAdded
		maxBagValue += maxValueAdded
		used[bestValueIdx] = true
	}

	return maxBagValue
}
