package main

import "fmt"

// What is the 1001st prime

// Target is the capacity of the list to find our prime.
const Target int = 10001

var counter int = 2
var primes []int = make([]int, 0)

func main() {
	for {
		if isPrime(counter) {
			primes = append(primes, counter)
		}

		if len(primes) == Target {
			break
		}

		counter++
	}

	fmt.Println(primes[Target-1])
}

func isPrime(x int) bool {
	for v := range primes {
		if x%primes[v] == 0 {
			return false
		}
	}

	for k := 11; k < x; k++ {
		if x%k == 0 {
			return false
		}
	}

	return true
}
