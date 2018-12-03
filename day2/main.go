package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func parseInput(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	result := make([]string, 0)
	for _, l := range bytes.Split(dat, []byte{'\n'}) {
		result = append(result, string(l))
	}
	return result
}

// hasNLetter detects whether the boxID contains N of the same letter
func hasNLetter(boxID string, N int) bool {
	seen := make(map[rune]int)
	for _, c := range(boxID) {
		seen[c] += 1
	}
	for _, count := range seen {
		if count == N {
			return true
		}
	}
	return false
}

func main() {
	boxIDs := parseInput("input.txt")

	// Part 1
	hasTwo := 0
	hasThree := 0
	for _, ID := range boxIDs {
		if hasNLetter(ID, 2) {
			hasTwo++
		}
		if hasNLetter(ID, 3) {
			hasThree++
		}
	}
	fmt.Println(hasTwo * hasThree)
}
