package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(word string) bool {

	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("enter a word")

	scanner.Scan()

	text := scanner.Text()

	fmt.Println(check(text))
}
