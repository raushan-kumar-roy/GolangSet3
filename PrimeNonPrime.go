package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countNonPrimeOddNumbers(n int) int {
	count := 0
	for i := 1; i < n; i += 2 {
		if !isPrime(i) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countNonPrimeOddNumbers(10))
	fmt.Println(countNonPrimeOddNumbers(20))
	fmt.Println(countNonPrimeOddNumbers(100))
}
