package main

import (
	"fmt"
)

var collect []int

func main() {
	collect = make([]int, 0, 10)
	nextInSequence(1, 2)
	fmt.Println(collect)
	evenNumbersOnly()
	fmt.Println(collect)

	var sum int
	for _, v := range collect {
		sum += v
	}
	fmt.Println(sum)
}

func nextInSequence(first int, second int) {
	nis := first + second
	collect = append(collect, nis)
	if nis < 4000000 {
		nextInSequence(second, nis)
	}
}

func evenNumbersOnly() {
	for i := 0; i < len(collect); i++ {
		fib := collect[i]
		if fib%2 != 0 {
			//fmt.Println("Removing", fib, "at index: ", i)
			collect = remove(fib)
			evenNumbersOnly()
		}
		//fmt.Println("Number is fine?", fib, "at index", i)
		//fmt.Println(fib, "%2==0:", (fib%2 == 0))
	}
}

func remove(value int) []int {
	a := make([]int, 0, 10)
	for _, v := range collect {
		if v != value {
			a = append(a, v)
		}
	}
	return a
}
