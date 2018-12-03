package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

func parseInput(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	result := make([]string, 0)
	for _, l := range bytes.Split(dat, []byte{'\n'}) {
		if len(l) == 0 { continue }
		result = append(result, string(l))
	}
	return result
}

// hasNLetter detects whether the boxID contains N of the same letter
func hasNLetter(boxID string, N int) bool {
	seen := make(map[rune]int)
	for _, c := range boxID {
		seen[c] += 1
	}
	for _, count := range seen {
		if count == N {
			return true
		}
	}
	return false
}

// return the count of different characters.
func diffCharCount(a, b string) int {
	if len(a) != len(b) {
		panic(fmt.Errorf("Could only compare strings with same length. %s, %s", a, b))
	}
	c := 0
	for i := range a {
		if a[i] != b[i] {
			c++
		}
	}
	return c
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

	// Part 2

	// Thoughts: ID has the same length, sort the list would put IDs diff
	// from each other one by one; so there is only need to compare adjacent
	// IDs in a sorted list.
	sort.Strings(boxIDs)
	for i, _ := range boxIDs {
		if i == 0 { continue }
		if diffCharCount(boxIDs[i - 1], boxIDs[i]) == 1 {
			fmt.Printf("%s\n%s\n", boxIDs[i-1], boxIDs[i])
		}

	}
}
