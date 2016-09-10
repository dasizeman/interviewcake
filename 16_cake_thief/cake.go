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

	maxValuesAtCapacities := make([]int, bagCapacity+1)

	for capacity := range maxValuesAtCapacities {

		maxValueAtCapacity := 0

		for _, cake := range cakeTypes {
			if cake.weight == 0 && cake.value > 0 {
				return math.MaxInt32
			}

			if cake.weight > capacity {
				continue
			}

			value := cake.value + maxValuesAtCapacities[capacity-cake.weight]
			if value > maxValueAtCapacity {
				maxValueAtCapacity = value
			}
		}

		maxValuesAtCapacities[capacity] = maxValueAtCapacity
	}
	return maxValuesAtCapacities[bagCapacity]
}
