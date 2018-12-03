package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func parseInput(filename string) []int {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(dat, []byte{'\n'})
	result := make([]int, 0)
	for _, l := range lines {
		if len(l) == 0 { continue }
		i, err := strconv.Atoi(string(l))
		if err != nil {
			fmt.Println("Error parsing: ", l, string(l))
			panic(err)
		}
		result = append(result, i)
	}
	return result
}

func part1() int {
	start := 0
	for _, i := range parseInput("input.txt") {
		start += i
	}
	return start
}

func part2() int {
	start := 0
	seen := make(map[int]bool)
	changes := parseInput("input.txt")

	// Note that your device might need to repeat its list of frequency
	// changes many times before a duplicate frequency is found.
Outer:
	for {
		for _, i := range changes {
			if seen[start] {
				break Outer
			}

			seen[start] = true
			start += i
		}
	}
	return start
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
