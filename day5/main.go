package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

func parseInput(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func IsReact(c1, c2 byte) bool {
	return int(math.Abs(float64(c1)-float64(c2))) == int('a')-int('A')
}

func DoReaction(in string) string {
	reacted := make(map[int]bool)
	for i := 0; i < len(in)-2;  {
		j := i + 1
		if IsReact(in[i], in[j]) {
			reacted[i] = true
			reacted[j] = true

			ii := i - 1
			jj := j + 1
			for {
				if ii <= 0 || jj >= len(in) - 1 {
					break
				}

				if IsReact(in[ii], in[jj]) {
					reacted[ii] = true
					reacted[jj] = true
				} else {
					break
				}
				ii--
				jj++
			}

			i = jj // the next unit which is not reacted
		} else {
			i++
		}
	}

	result := ""
	removed := ""

	for i := range in {
		if !reacted[i] {
			result += string(in[i])
		} else {
			removed += string(in[i])
		}
	}
	// fmt.Println(removed)
	return result
}

func main() {
	polymers := parseInput("input.txt")
	for {
		ol := len(polymers)
		polymers = DoReaction(polymers)
		l := len(polymers)
		if ol == l {
			break
		} else {
			fmt.Println(l, ol)
			// break
		}
	}
	fmt.Println(len(polymers))
}
