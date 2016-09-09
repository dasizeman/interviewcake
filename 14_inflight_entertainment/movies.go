package main

import (
	"fmt"
)

func main() {
	moviesTimesGood := []int{125, 87, 200, 150, 50}
	moviesTimesBad := []int{125, 87, 200, 150}

	fmt.Printf("Found movies (should be true): %t\n",
		twoMoviesForTheFlight(250, moviesTimesGood))

	fmt.Printf("Found movies (should be false): %t\n",
		twoMoviesForTheFlight(250, moviesTimesBad))
}

func twoMoviesForTheFlight(flightTime int, movieLengths []int) bool {
	// This map will have the format remainders[remainingtime] = index of movie
	//used for this remainder
	remainders := make(map[int]int)

	// Build the remainder map
	for idx, time := range movieLengths {
		remainders[flightTime-time] = idx
	}

	// Find if any of the movies fit a remainder time
	for idx, time := range movieLengths {
		otherIdx, found := remainders[time]

		if found && otherIdx != idx {
			return true
		}
	}

	return false
}
