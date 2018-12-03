package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Claim struct {
	ID         int
	LeftOffset int
	TopOffset  int
	Width      int
	Height     int
}

func (c *Claim) GetRightOffset() int {
	return c.LeftOffset + c.Width
}

func (c *Claim) GetBottomOffset() int {
	return c.TopOffset + c.Height
}

func (c *Claim) String() string {
	return fmt.Sprintf("Claim<ID: %d, %d, %d, %d, %d>", c.ID, c.LeftOffset, c.TopOffset, c.Width, c.Height)
}

func mustInt(i int, err error) int {
	if err != nil {
		panic(err)
	}
	return i
}

func parseInput(filename string) []Claim {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	result := make([]Claim, 0)
	for _, l := range bytes.Split(dat, []byte{'\n'}) {
		if len(l) == 0 {
			continue
		}
		// parse the line in the format: #1 @ 596,731: 11x27
		l := bytes.Split(l, []byte{'@'})

		i := mustInt(strconv.Atoi(
			strings.Trim(string(l[0][1:]), " "))) // skip '#'

		l = bytes.Split(l[1], []byte{':'})
		offsets := bytes.Split(l[0], []byte{','})
		leftOff := mustInt(strconv.Atoi(
			strings.Trim(string(offsets[0]), " ")))
		topOff := mustInt(strconv.Atoi(
			strings.Trim(string(offsets[1]), " ")))

		dims := bytes.Split(l[1], []byte{'x'})
		width := mustInt(strconv.Atoi(
			strings.Trim(string(dims[0]), " ")))
		height := mustInt(strconv.Atoi(
			strings.Trim(string(dims[1]), " ")))

		claim := Claim{
			ID:         i,
			LeftOffset: leftOff,
			TopOffset:  topOff,
			Width:      width,
			Height:     height,
		}
		result = append(result, claim)

	}
	return result
}

type Fabric struct {
	Width  int
	Height int
	Grid   [][][]*Claim
}

func InitFabric(width, height int) *Fabric {
	f := Fabric{Width: width, Height: height}
	f.Grid = make([][][]*Claim, height)
	for i := range f.Grid {
		f.Grid[i] = make([][]*Claim, width)
		for j := range f.Grid[i] {
			f.Grid[i][j] = make([]*Claim, 0)
		}
	}
	return &f
}

func (f *Fabric) AddClaim(claim Claim) {
	for j := claim.TopOffset; j < claim.GetBottomOffset(); j++ {
		for i := claim.LeftOffset; i < claim.GetRightOffset(); i++ {
			f.Grid[j][i] = append(f.Grid[j][i], &claim)
		}
	}
}

func (f *Fabric) FindOverlapInches() int {
	count := 0
	for j := 0; j < f.Height; j++ {
		for i := 0; i < f.Width; i++ {
			if len(f.Grid[j][i]) >= 2 {
				count += 1
			}
		}
	}
	return count
}

// getFabricDim returns the dimension for the fabric to contain all the claims.
// This is to ensure the cases where fabric needed are larger than 1000 inches.
func getFabricDim(claims []Claim) (int, int) {
	rightMost := 0
	bottomMost := 0
	for _, c := range claims {
		if c.GetBottomOffset() > bottomMost {
			bottomMost = c.GetBottomOffset()
		}

		if c.GetRightOffset() > rightMost {
			rightMost = c.GetRightOffset()
		}
	}
	return rightMost, bottomMost
}

func main() {
	claims := parseInput("input.txt")
	width, height := getFabricDim(claims)

	// Option 1: Keep the fabric in memory, and while iterating through
	// claims, save the claim per inch^2.
	fabric := InitFabric(width, height)
	for _, c := range claims {
		fabric.AddClaim(c)
	}
	fmt.Println(fabric.FindOverlapInches())
}
