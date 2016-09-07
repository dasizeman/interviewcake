package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Enter some sample stock values:")
	scanner := bufio.NewScanner(os.Stdin)

	var stockPrices []int

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Please enter an integer\n")
			continue
		}
		stockPrices = append(stockPrices, num)
	}

	fmt.Printf("You entered: %v\n", stockPrices)

	fmt.Printf("Maximum profit is: %d", getMaxProfit(stockPrices))

}

func getMaxProfit(stockPricesYesterday []int) (maxProfit int) {
	if len(stockPricesYesterday) < 2 {
		log.Fatalf("Need at least two stock samples\n")
	}

	currentLowest := stockPricesYesterday[0]
	maxProfit = stockPricesYesterday[1] - currentLowest

	for _, price := range stockPricesYesterday[1:] {

		// See if we could make a better profit buying at the current lowest price and selling now
		potentialProfit := price - currentLowest
		if potentialProfit > maxProfit {
			maxProfit = potentialProfit
		}
		// Make sure the minimum price is up to date
		if price < currentLowest {
			currentLowest = price
		}

	}

	return
}
