package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func wordAmountCount(texts string) map[string]int {
	texts = strings.ToLower(texts)
	a := regexp.MustCompile(`[^\w\s]`)

	cleanWords := a.ReplaceAllString(texts, "")

	words := strings.Fields(cleanWords)
	wordFrequency := make(map[string]int)
	for _, word := range words {
		wordFrequency[word]++
	}

	return wordFrequency

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("write your texts")
	scanner.Scan()

	texts := scanner.Text()

	fmt.Println(wordAmountCount(texts))
}
