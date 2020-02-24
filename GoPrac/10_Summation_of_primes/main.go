package main

import (
	"fmt"
)

const Limit int = 2000000

func sumPrimes(num int, prime []bool) {
	total := 0
	for i := 2; i <= num; i++ {
		if prime[i] == false {
			total += i
		}
	}

	fmt.Printf("Sum of all primes less than %d is %d\n", num, total)
}

func sieve(num int) {
	prime := make([]bool, num+1)
	for i := 0; i < num+1; i++ {
		prime[i] = false
	}

	// this is that iteration upto the square root of the limit number
	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true //cross
			}
		}
	}

	sumPrimes(num, prime)
}

func main() {
	sieve(Limit)
}
