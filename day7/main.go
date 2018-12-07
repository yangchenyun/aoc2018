package main

import (
	"fmt"
	"os"
	"sort"
)

type Edge struct {
	From string
	To   string
}

type Node struct {
	Name     string
	Depends  []*Node
	Nexts    []*Node
	Seconds  int
	FinishOn int
}

func (n *Node) String() string {
	depends := make([]string, 0)
	nexts := make([]string, 0)
	for _, n := range n.Depends {
		depends = append(depends, n.Name)
	}

	for _, n := range n.Nexts {
		nexts = append(nexts, n.Name)
	}
	// return fmt.Sprintf("<%s depends=%s, nexts=%v, finish=%d>", n.Name, depends, nexts, n.FinishOn)
	return fmt.Sprintf("<%s, %d>", n.Name, n.FinishOn)
}

func parseInput() []*Node {
	f, _ := os.Open("input.txt")
	nodeMap := make(map[string]*Node)
	for {
		var a, b string
		_, err := fmt.Fscanf(f, "Step %s must be finished before step %s can begin.\n", &a, &b)
		if err != nil {
			break
		}
		if _, ok := nodeMap[a]; !ok {
			nodeMap[a] = &Node{a, make([]*Node, 0), make([]*Node, 0), 60 + int(a[0]-'A'+1), 0}
		}
		if _, ok := nodeMap[b]; !ok {
			nodeMap[b] = &Node{b, make([]*Node, 0), make([]*Node, 0), 60 + int(b[0]-'A'+1), 0}
		}
		nodeA := nodeMap[a]
		nodeB := nodeMap[b]
		nodeB.Depends = append(nodeB.Depends, nodeMap[a])
		nodeA.Nexts = append(nodeA.Nexts, nodeMap[b])
	}
	result := make([]*Node, 0)
	for _, node := range nodeMap {
		result = append(result, node)
	}
	return result
}

func FindBegins(nodes []*Node) []*Node {
	begins := make([]*Node, 0)
	for _, node := range nodes {
		if len(node.Depends) == 0 {
			begins = append(begins, node)
		}
	}
	return begins
}

func FindEnd(nodes []*Node) *Node {
	for _, node := range nodes {
		if len(node.Nexts) == 0 {
			return node
		}
	}
	return nil
}

func Contains(n *Node, nodes []*Node) bool {
	for _, node := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

func IsSubset(sub, sup []*Node) bool {
	for _, nn := range sub {
		if !Contains(nn, sup) {
			return false
		}
	}
	return true
}

func FindNextIdx(nodes []*Node, finished []*Node) int {
	for i, n := range nodes {
		if IsSubset(n.Depends, finished) {
			return i
		}
	}
	return -1
}

// AssignWork at the time.
func AssignWork(workers []int, now int, finish int) bool {
	for i := range workers {
		if workers[i] <= now {
			workers[i] = finish
			return true
		}
	}
	return false
}

func WorkersAvailable(workers []int, time int) bool {
	for i := range workers {
		if workers[i] <= time {
			return true
		}
	}
	return false
}

func main() {
	nodes := parseInput()
	begins := FindBegins(nodes)
	end := FindEnd(nodes)

	// Part 1
	nexts := make([]*Node, len(begins))
	copy(nexts, begins)
	finished := make([]*Node, 0)
	steps := ""
	for {
		if Contains(end, nexts) && IsSubset(end.Depends, finished) {
			steps += end.Name
			break
		}

		sort.SliceStable(nexts, func(i, j int) bool {
			return nexts[i].Name < nexts[j].Name
		})
		i := FindNextIdx(nexts, finished)
		next := nexts[i]
		nexts = append(nexts[:i], nexts[i+1:]...)

		// Add new nodes and sort
		for _, n := range next.Nexts {
			if !Contains(n, finished) && !Contains(n, nexts) {
				nexts = append(nexts, n)
			}
		}

		finished = append(finished, next)
		steps += next.Name
	}
	fmt.Println(steps)

	// Part 2
	now := 0
	workers := make([]int, 5) // Each int represents the time in the future when the work is available.
	working := make([]*Node, 0)
	nexts = make([]*Node, len(begins))
	copy(nexts, begins)
	finished = make([]*Node, 0)
	for {
		// stage to assign work for the moment
		for {
			if !WorkersAvailable(workers, now) {
				break
			}

			sort.SliceStable(nexts, func(i, j int) bool {
				return nexts[i].Name < nexts[j].Name
			})
			i := FindNextIdx(nexts, finished)
			if i == -1 {
				break  // when no work is needed
			}

			next := nexts[i]
			nexts = append(nexts[:i], nexts[i+1:]...)
			working = append(working, next)
			next.FinishOn = now + next.Seconds
			AssignWork(workers, now, now+next.Seconds)
		}

		// check work completion
		sort.SliceStable(working, func(i, j int) bool {
			return working[i].FinishOn < working[j].FinishOn
		})
		work := working[0]
		working = working[1:]
		for _, n := range work.Nexts {
			if !Contains(n, finished) && !Contains(n, nexts) {
				nexts = append(nexts, n)
			}
		}
		finished = append(finished, work)
		now = work.FinishOn

		if len(nexts) == 0 {
			break
		}
	}
	fmt.Println(now)
}
