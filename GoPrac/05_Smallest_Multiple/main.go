package main

import "fmt"

/**
Smallest number divisible by all numbers from 1 to 20
*/
func main() {
	counter := 2520 //given as divisible by numbers 1 to 10

	for i := 20; i > 0; i-- {
		if counter%i == 0 {
			continue
		}

		counter++
		i = 20
	}
	fmt.Println(counter)
}
