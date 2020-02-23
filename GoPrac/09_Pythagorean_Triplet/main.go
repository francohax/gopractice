package main

import (
	"math"
)

const pathTotal int = 1000

func main() {
}

func isNaturalNumber(i float64) bool {
	return i == math.Trunc(i)
}
