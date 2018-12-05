package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/golang-collections/collections/stack"
)

const caseDiff = byte('a' - 'A')

func parseInput(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(dat), "\n")
}

func IsReact(c1, c2 byte) bool {
	return c1 - caseDiff == c2 || c1 + caseDiff == c2
}

func DoReaction(in string) string {
	stack := stack.New()
	for i := range in {
		if stack.Len() > 0 && IsReact(stack.Peek().(byte), in[i]) {
			stack.Pop()
		} else {
			stack.Push(in[i])
		}
	}
	result := ""
	for {
		if stack.Len() == 0 {
			break
		}
		result += string(stack.Pop().(byte))
	}
	return result
}

// WithoutUnit produces a new polymer without unit.
func WithoutUnit(polymers string, unit byte) string {
	result := ""
	for i := range polymers {
		if IsReact(polymers[i], unit) || polymers[i] == unit {
			continue
		}
		result += string(polymers[i])
	}
	return result
}

func FindUnits(polymers string) []byte {
	seen := make(map[byte]bool)
	polymers = strings.ToLower(polymers)
	for i := range polymers {
		seen[polymers[i]] = true
	}
	result := make([]byte, 0)
	for k := range seen {
		result = append(result, k)
	}
	return result
}

func main() {
	polymers := parseInput("input.txt")
	// part 1
	fmt.Println(len(DoReaction(polymers)))

	// part 2
	lens := make([]int, 0)
	for _, unit := range FindUnits(polymers) {
		p := WithoutUnit(polymers, unit)
		lens = append(lens, len(DoReaction(p)))
	}
	sort.Ints(lens)
	fmt.Println(lens[0])
}
