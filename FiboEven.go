package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func countEvenValues(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		if fibonacci(i)%2 == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countEvenValues(10))
}
