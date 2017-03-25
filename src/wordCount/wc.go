package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
//
// You might find strings.Fields helpful.

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	myMap := make(map[string]int)
	for _, word := range words {
		myMap[word] += 1
	}
	return myMap
}

func main() {
	wc.Test(WordCount)
}
