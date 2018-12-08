package main

import (
	"fmt"
	"os"
)

type Node struct {
	MetaNum   int
	Metas     []int
	MetaValue int

	ChildNum int
	Children []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("[meta_num=%d meta=%d child_num=%d children=%s]", n.MetaNum, n.Metas, n.ChildNum, n.Children)
}

// PraseTree returns the root node of a tree and the remaining input
func ParseTree(input []int) (*Node, []int) {
	// construct node
	n := &Node{}
	n.ChildNum = input[0]
	n.Children = make([]*Node, input[0])
	n.MetaNum = input[1]
	n.Metas = make([]int, input[1])

	// parse children
	remaining := input[2:]
	var child *Node
	for i := range n.Children {
		child, remaining = ParseTree(remaining)
		n.Children[i] = child
	}

	// parse metadata
	for i := range n.Metas {
		n.Metas[i] = remaining[i]
	}
	remaining = remaining[len(n.Metas):]
	return n, remaining
}

// Traverse the tree
func Traverse(root *Node, post func(n *Node)) {
	for _, child := range root.Children {
		Traverse(child, post)
	}
	post(root)
}

func main() {
	// parse
	f, _ := os.Open("input.txt")
	result := make([]int, 0)
	for {
		var i int
		_, err := fmt.Fscan(f, &i)
		if err != nil {
			break
		}
		result = append(result, i)
	}
	// result = []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}
	root, _ := ParseTree(result)

	// part 1
	metaSum := 0
	post := func(n *Node) {
		for _, m := range n.Metas {
			metaSum += m
		}
	}
	Traverse(root, post)
	fmt.Println(metaSum)

	// part 2
	post = func(n *Node) {
		var metaVal int
		if n.ChildNum == 0 {
			for _, m := range n.Metas {
				metaVal += m
			}
		} else {
			for _, m := range n.Metas {
				i := m - 1
				if 0 <= i && i < n.ChildNum {  // check index is valid
					child := n.Children[i]
					metaVal += child.MetaValue
				}
			}
		}
		n.MetaValue = metaVal
	}
	Traverse(root, post)
	fmt.Println(root.MetaValue)
}
