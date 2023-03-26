package main

import (
	"fmt"
	"math"
)

type treenode struct {
	data  int
	left  *treenode
	right *treenode
}

func (t *treenode) BST() bool {
	min := math.MinInt
	max := math.MaxInt
	return t.BSTHelper(min, max)
}
func (t *treenode) BSTHelper(min, max int) bool {
	if t == nil {
		return true
	}
	if t.data < min || t.data > max {
		return false
	}
	return t.left.BSTHelper(min, t.data-1) && t.right.BSTHelper(t.data+1, max)
}
func main() {
	var tn treenode
	tn.data = 15
	tn.left = &treenode{data: 12}
	tn.right = &treenode{data: 18}
	tn.left.left = &treenode{data: 10}
	tn.left.right = &treenode{data: 13}
	tn.right.left = &treenode{data: 17}
	tn.right.right = &treenode{data: 20}
	if tn.BST() {
		fmt.Println("Tree is a BST")
	} else {
		fmt.Println("Not a BST")
	}
}
