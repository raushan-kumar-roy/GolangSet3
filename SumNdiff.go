package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func getSum_n_Diff(a int, b int) (int, int) {
	sum := sum(a, b)
	diff := diff(a, b)
	return sum, diff
}

func main() {
	var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	sum, diff := getSum_n_Diff(a, b)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
}
