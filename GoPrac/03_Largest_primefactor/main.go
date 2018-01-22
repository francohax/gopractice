package main

func main() {
	x := 600851475143
	collection := make([]int, 0, 10)
	for i := 3; i < x; i++ {
		//fmt.Println(i)
		//fmt.Println("divisible", x%i)
		if isPrime(i) {
			collection = append(collection, i)
		}
	}
}

func reachesTarget(factors []int) bool {
	return false
}

func isPrime(x int) bool {
	for k := 2; k < x; k++ {
		if x%k == 0 {
			return false
		}
	}
	return true
}
