package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func parseInput(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(dat), "\n")
}

func IsReact(c1, c2 byte) bool {
	return int(math.Abs(float64(c1)-float64(c2))) == int('a')-int('A')
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

func main() {
	polymers := parseInput("input.txt")
        // part 1
        fmt.Println(len(DoReaction(polymers)))
	}
	fmt.Println(len(polymers))
}
