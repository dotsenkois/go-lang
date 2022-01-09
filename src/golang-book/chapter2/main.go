package main

import (
	"fmt"
)

func fibonacchi(x int) int {
	if x < 2 {
		return 1
	}
	return fibonacchi(x-2) + fibonacchi(x-1)
}

func printFib(x int) {
	for i := 1; i < x; i++ {
		fmt.Print(fibonacchi(i), " ")
	}

}
func main() {
	printFib(10)
}
