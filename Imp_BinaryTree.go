package main

import (
	"fmt"
	"math"
)

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

func display(t *treeNode) {
	if t == nil {
		return
	}
	display(t.left)
	fmt.Print(t.data, "-")
	display(t.right)
}

func (t *treeNode) getTreeNodeCount() int {
	if t == nil {
		return 0
	}
	return t.left.getTreeNodeCount() + t.right.getTreeNodeCount() + 1
}
func (t *treeNode) TreeDegree() int {
	degree := 0
	if t == nil {
		return degree
	}
	left := t.left.TreeDegree()
	right := t.right.TreeDegree()
	return int(math.Max(float64(left), float64(right))) + 1
}
func main() {
	root := &treeNode{data: 10}
	root.left = &treeNode{data: 9}
	root.right = &treeNode{data: 15}
	root.left.left = &treeNode{data: 20}
	root.left.right = &treeNode{data: 25}
	root.right.left = &treeNode{data: 12}
	root.right.right = &treeNode{data: 30}
	root.left.left.right = &treeNode{data: 50}
	root.left.left.right.left = &treeNode{data: 100}
	display(root)
	fmt.Println("")
	count := root.getTreeNodeCount()
	fmt.Println("Number of nodes:", count)
	depth := root.TreeDegree()
	fmt.Println("height of tree:", depth)
}
