package main

import "fmt"

// Difference between the sum of squares and square of sums of a number (100).
func main() {
	x := 100
	fmt.Println(squareOfSum(x) - sumOfSquares(x))
}

func sumOfSquares(x int) int {
	sum := 0
	for i := 0; i <= x; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum(x int) int {
	sum := 0
	for i := 0; i <= x; i++ {
		sum += i
	}
	return sum * sum
}
