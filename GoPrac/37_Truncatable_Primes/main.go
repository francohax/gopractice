package main

import "log"
import "strconv"
import "strings"

func main() {
	log.Println("Search Started")
	digits := make([]int, 0, 15)
	sums := make([]int, 0, 15)
	digits, sums = loadPrimes(15);
	log.Println("Collection of valid digits: ", digits)
	log.Println("Sum of the truncated primes of the valid digits: ", sums)
	log.Println("End of Service")
}

func loadPrimes(capacity int) ([]int, []int) {
	digits := make([]int, 0, 15)
	sumOfPrimes := make([]int, 0, 15)

	truncatedPrimes := make([][]int, 0, 15)
	currentNumber := 1000

	for (len(sumOfPrimes) < capacity) {
		currentNumber++
		num := strconv.Itoa(currentNumber)
		log.Println("Current Digit: ", num)

		if (!strings.Contains(num, "0")) {
			successL, resL := truncateLeft(num)
			if successL {
				log.Println("Digit Successfully Truncated LEFT: ", resL)
				truncatedPrimes = append(truncatedPrimes, resL)
			} else { continue }

			successR, resR := truncateRight(num)
			if successR {
				log.Println("Digit Successfully Truncated RIGHT: ", resR)
				truncatedPrimes = append(truncatedPrimes, resR)
			} else { continue }

			var sumOfPrime int
			for _, primes := range truncatedPrimes {
				for _, prime := range primes {
					sumOfPrime += prime
				}
			}

			digits = append(digits, currentNumber)
			log.Println("Adding to Digit: ", digits)
			sumOfPrimes = append(sumOfPrimes, sumOfPrime)
			log.Println("Adding to Sum Of Primes: ", sumOfPrimes)
		}
	}

	return digits, sumOfPrimes
}

// incrementally truncating the string (left to right)
func truncateLeft(t string) (bool, []int) {
	truncCache := make([]int, 0, 5)
	for d := 0; d < len(t); d++ {
		ok := truncateValue(t, d, len(t), &truncCache)
		
		if !ok {
			return false, truncCache
		}
	}

	//log.Println("Appended to LEFT cache:", truncCache)
	return true, truncCache
}

// decrementally truncating the string (right to left)
func truncateRight(t string) (bool, []int) {
	truncCache := make([]int, 0, 5)
	for d := 0; d < len(t); d++ {
		ok := truncateValue(t, 0, len(t)-d, &truncCache)

		if !ok {
			return false, truncCache
		}
	}

	//log.Println("Appended to RIGHT cache:", truncCache)
	return true, truncCache
}

func truncateValue(fullValue string, start int, end int, cache *[]int) bool {
	trunc, err := strconv.Atoi(fullValue[start: end])
	//log.Println("Truncated Value: ", trunc)
	if err != nil {
		panic(err)
	}

	if !validateTruncValue(trunc) {
		return false
	}

	*cache = append(*cache, trunc)
	return true
}

//TO DO
func validateTruncValue(trunc int) bool {
	if !isPrime(trunc) {
		//log.Println(trunc, " is NOT prime")
		return false
	} 

	//log.Println(trunc, " is prime")
	return true
}

func isPrime(x int) bool {
	for k := 2; k < x; k++ {
		if x%k == 0 {
			return false
		}
	}
	return true
}