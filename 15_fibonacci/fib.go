package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("f(%d)=%d\n", i, fib(i))
	}
}

func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	nMinusOne := 1
	nMinusTwo := 0
	for i := 2; i <= n; i++ {
		newNum := nMinusTwo + nMinusOne
		nMinusTwo = nMinusOne
		nMinusOne = newNum
	}

	return nMinusOne
}
