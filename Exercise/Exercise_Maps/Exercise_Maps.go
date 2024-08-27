package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	maps := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		maps[word] = 1
	}
	return maps
}

func main() {
	wc.Test(WordCount)
}