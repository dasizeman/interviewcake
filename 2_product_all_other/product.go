package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var amt, max int
	fmt.Printf("How many values?\n")
	_, err := fmt.Scanf("%d\n", &amt)
	fmt.Printf("Max value?\n")
	_, err = fmt.Scanf("%d\n", &max)

	values := randomSlice(amt, max)

	fmt.Printf("Generated: %v\n", values)

	if err != nil {
		fmt.Printf("Something wrong with that input.\n")
	}

	// interviewcake test case
	//values := []int{1, 7, 3, 4}

	result := getProductsOfAllIntsExceptAtIndex(values)
	fmt.Printf("Product result: %v\n", result)
}

func randomSlice(len, max int) (res []int) {
	for i := 0; i < len; i++ {
		res = append(res, rand.Intn(max)+1)
	}
	return res
}

func getProductsOfAllIntsExceptAtIndex(a []int) (res []int) {
	// First we run through the input and create a data structure that holds the product of all values up to that index, one for each direction
	products := make([]int, len(a))
	products[0] = 1

	currentProduct := 1

	// In the following loops we compute the result values by greedily finding the product of all values up to and not including the value at the current index in both directions, and multiplying these values together
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
