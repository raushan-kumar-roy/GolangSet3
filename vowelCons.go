package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var vowel, consonent int
	vowel = 0
	consonent = 0

	fmt.Println("Enter sentence:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentense := scanner.Text()
	fmt.Println(sentense)

	for i := 0; i < len(sentense); i++ {
		if sentense[i] == ' ' {
			continue
		} else if sentense[i] == 'a' || sentense[i] == 'e' || sentense[i] == 'i' || sentense[i] == 'o' || sentense[i] == 'u' {
			vowel += 1
		} else {
			consonent += 1
		}
	}
	fmt.Println("Vowel : ", vowel)
	fmt.Println("Consonent : ", consonent)

}
