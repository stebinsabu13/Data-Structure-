package main

import "fmt"

type treeNode struct {
	data     int
	children []*Tree
}
type Tree struct {
	root *treeNode
}
type queue struct {
	arr []*Tree
}

func (q *queue) enqueue(node *Tree) {
	q.arr = append(q.arr, node)
}
func (q *queue) dequeue() {
	q.arr = q.arr[1:]
}

func (t *Tree) insert(node Tree) {
	if t.root == nil {
		t.root = node.root
		return
	}
	t.root.children = append(t.root.children, &node)
}
func (t *Tree) display() {
	var q queue
	if t.root == nil {
		return
	}
	q.enqueue(t)
	for len(q.arr) != 0 {
		n := len(q.arr)
		for ; n > 0; n-- {
			node := q.arr[0]
			fmt.Print(node.root.data, "-")
			q.dequeue()
			for i := 0; i < len(node.root.children); i++ {
				q.enqueue(node.root.children[i])
			}
		}
		fmt.Println()
	}
}
func main() {
	var t Tree
	t.insert(Tree{root: &treeNode{data: 10}})
	t.insert(Tree{root: &treeNode{data: 20}})
	t.insert(Tree{root: &treeNode{data: 30}})
	t.root.children[1].insert(Tree{root: &treeNode{data: 40}})
	t.root.children[1].insert(Tree{root: &treeNode{data: 140}})
	t.root.children[1].root.children[0].insert(Tree{root: &treeNode{data: 240}})
	t.insert(Tree{root: &treeNode{data: 50}})
	t.insert(Tree{root: &treeNode{data: 60}})
	t.root.children[2].insert(Tree{root: &treeNode{data: 400}})
	t.insert(Tree{root: &treeNode{data: 70}})
	t.insert(Tree{root: &treeNode{data: 80}})
	t.display()
}
