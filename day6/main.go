package main

import (
	"fmt"
	"os"
	"sort"
)

type ID int
type Coord struct {
	ID ID
	X  int64
	Y  int64
}

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func parseInput(fname string) []Coord {
	result := make([]Coord, 0)

	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	var id ID
	for {
		c := Coord{ID: id}
		_, err := fmt.Fscanf(f, "%d, %d", &c.X, &c.Y)
		if err != nil {
			break
		}
		result = append(result, c)
		id++
	}
	return result
}

func ManhattanDistance(c1, c2 Coord) int {
	return int(abs(c1.X-c2.X) + abs(c1.Y-c2.Y))
}

type View struct {
	SX int64
	SY int64
	EX int64
	EY int64
	grid [][]ID
}

func InitView(coords []Coord) *View {
	sort.SliceStable(coords, func(i, j int) bool {
		return coords[i].X > coords[j].X
	})
	maxX := coords[0].X
	minX := coords[len(coords)-1].X

	sort.SliceStable(coords, func(i, j int) bool {
		return coords[i].Y > coords[j].Y
	})
	maxY := coords[0].Y
	minY := coords[len(coords)-1].Y


	view := &View{minX, minY, maxX, maxY, nil}
	grid := make([][]ID, view.Height())
	for i := range grid {
		grid[i] = make([]ID, view.Width())
	}
	view.grid = grid
	return view
}

func (v *View) Width() int {
	return int(v.EX - v.SX)
}

func (v *View) Height() int {
	return int(v.EY - v.SY)
}

func (v *View) EdgeIDs() map[ID]bool {
	edgeIDs := make(map[ID]bool)
	for i := 0; i < v.Height(); i++ {
		edgeIDs[v.grid[i][0]] = true
		edgeIDs[v.grid[i][v.Width()-1]] = true
	}

	for j := 0; j < v.Width(); j++ {
		edgeIDs[v.grid[0][j]] = true
		edgeIDs[v.grid[v.Height()-1][j]] = true
	}
	return edgeIDs
}

func (v *View) AreaMap() map[ID]int {

	areaMap := make(map[ID]int)
	for i := 0; i < v.Height(); i++ {
		for j := 0; j < v.Width(); j++ {
			if v.grid[i][j] != -1 {
				areaMap[v.grid[i][j]]++
			}
		}
	}
	return areaMap
}

// FindCloestCoord finds the closest one, if there are more than two, return none
func FindCloestCoord(c Coord, coords []Coord) *Coord {
	dmap := make(map[Coord]int)
	dlist := make([]int, 0)

	for _, cc := range coords {
		d := ManhattanDistance(c, cc)
		dmap[cc] = d
		dlist = append(dlist, d)
	}
	sort.Ints(dlist)

	// When c is equally close to two coordinators
	if dlist[0] == dlist[1] {
		return nil
	}
	minD := dlist[0]
	for k, d := range dmap {
		if d == minD {
			return &k
		}
	}
	return nil
}

func FillDistances(view *View, coords []Coord) {
	for i := 0; i < view.Height(); i++ {
		for j := 0; j < view.Width(); j++ {
			gridCoord := Coord{
				X: view.SX + int64(j),
				Y: view.SY + int64(i),
			}
			cc := FindCloestCoord(gridCoord, coords)
			if cc != nil {
				view.grid[i][j] = cc.ID
			} else {
				view.grid[i][j] = -1
			}
		}
	}
}

func TotalDistance(c Coord, coords []Coord) int {
	result := 0
	for _, cc := range coords {
		result += ManhattanDistance(c, cc)
	}
	return result
}

const IsRegion = 1
func FillRegion(view *View, coords []Coord, maxD int) {
	for i := 0; i < view.Height(); i++ {
		for j := 0; j < view.Width(); j++ {
			gridCoord := Coord{
				X: view.SX + int64(j),
				Y: view.SY + int64(i),
			}
			d := TotalDistance(gridCoord, coords)
			if d < maxD {
				view.grid[i][j] = IsRegion
			}
		}
	}
}

func RegionSize(view *View) int {
	result := 0
	for i := 0; i < view.Height(); i++ {
		for j := 0; j < view.Width(); j++ {
			if view.grid[i][j] == IsRegion {
				result++
			}
		}
	}
	return result
}


func main() {
	coords := parseInput("input.txt")

	// part 1
	view := InitView(coords)
	FillDistances(view, coords)
	areaMap := view.AreaMap()

	areas := make([]int, len(areaMap))
	for id, area := range areaMap {
		edgeIDs := view.EdgeIDs()
		if !edgeIDs[id] {
			areas = append(areas, area)
		}
	}
	sort.Ints(areas)
	fmt.Println(areas[len(areas)-1])

	// // part 2
	view = InitView(coords)
	FillRegion(view, coords, 10000)
	fmt.Println(RegionSize(view))

}
