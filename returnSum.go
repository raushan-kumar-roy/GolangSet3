package main

import "fmt"

func sumMultiples(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
func main() {
	result := sumMultiples(10)
	fmt.Println(result)

}
