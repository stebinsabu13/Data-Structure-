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
type Tree struct {
	root *treeNode
}

func (t *Tree) insert(val int) {
	if t.root == nil {
		t.root = &treeNode{data: val}
	} else {
		t.root.Insert(val)
	}
}
func (n *treeNode) Insert(val int) {
	if val < n.data {
		if n.left == nil {
			n.left = &treeNode{data: val}
		} else {
			n.left.Insert(val)
		}
	} else if val > n.data {
		if n.right == nil {
			n.right = &treeNode{data: val}
		} else {
			n.right.Insert(val)
		}
	}
}

func (t *Tree) delete(val int) {
	if t.root == nil {
		fmt.Println("Tree empty")
	} else {
		if t.root.Search(val) {
			t.root.Delete(val)
		} else {
			fmt.Println("Value doesn't exsists")
		}
	}
}
func (n *treeNode) Delete(val int) *treeNode {
	if val < n.data {
		n.left = n.left.Delete(val)
	} else if val > n.data {
		n.right = n.right.Delete(val)
	} else {
		if n.left == nil && n.right == nil {
			return nil
		} else if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		} else {
			node := findSuccessor(n.right)
			n.data = node.data
			n.right = n.right.Delete(node.data)
		}
	}
	return n
}
func findSuccessor(s *treeNode) *treeNode {
	for s.left != nil {
		s = s.left
	}
	return s
}
func trimBST(root *treeNode, low int, high int) *treeNode {
	if root == nil {
		return nil
	}
	if root.data < low {
		return trimBST(root.right, low, high)
	}
	if root.data > high {
		return trimBST(root.left, low, high)
	}
	root.left = trimBST(root.left, low, high)
	root.right = trimBST(root.right, low, high)
	return root
}
func (t *Tree) search(val int) bool {
	if t == nil {
		return false
	} else {
		return t.root.Search(val)
	}
}
func (n *treeNode) Search(val int) bool {
	if n == nil {
		return false
	}
	if val > n.data {
		return n.right.Search(val)
	} else if val < n.data {
		return n.left.Search(val)
	} else {
		return true
	}
}
func levelNodes(n *treeNode, level int) {
	if n == nil {
		return
	}
	if level == 1 {
		fmt.Print(n.data, "-")
	} else {
		levelNodes(n.left, level-1)
		levelNodes(n.right, level-1)
	}
}
func (t *Tree) leafnodes() {
	if t == nil {
		fmt.Println("Tree is empty")
	} else {
		t.root.leafnodes()
	}
}
func (n *treeNode) leafnodes() {
	if n.left == nil && n.right == nil {
		fmt.Print(n.data, "-")
		return
	}
	if n.left != nil {
		n.left.leafnodes()
	}
	if n.right != nil {
		n.right.leafnodes()
	}
}
func (t *Tree) nonleafnodes() {
	if t == nil {
		fmt.Println("Tree is empty")
	} else {
		t.root.nonleafnodes()
	}
}
func (n *treeNode) nonleafnodes() {
	if n.left != nil || n.right != nil {
		fmt.Print(n.data, "-")
		if n.left == nil {
			n.right.nonleafnodes()
		} else if n.right == nil {
			n.left.nonleafnodes()
		} else {
			n.left.nonleafnodes()
			n.right.nonleafnodes()
		}
	}
}
func (t *Tree) rightview() {
	if t == nil {
		fmt.Println("Tree is empty")
	} else {
		t.root.rightview()
	}
}
func (n *treeNode) rightview() {
	if n.right != nil {
		fmt.Print(n.data, "-")
		n.right.rightview()
	} else {
		fmt.Print(n.data, "-")
	}
}
func (t *Tree) leftview() {
	if t == nil {
		fmt.Println("Tree is empty")
	} else {
		t.root.leftview()
	}
}
func (n *treeNode) leftview() {
	if n.left != nil {
		fmt.Print(n.data, "-")
		n.left.leftview()
	} else {
		fmt.Print(n.data, "-")
	}
}
func (t *Tree) closetValue(val int) {
	if t == nil {
		fmt.Println("Tree is empty")
	} else {
		fmt.Println(t.root.closetValue(val))
	}
}
func (n *treeNode) closetValue(val int) int {
	diff := math.MaxInt
	var cudiff int
	var cudata int
	for n != nil {
		if n.data < val {
			cudiff = val - n.data
			if cudiff < diff {
				diff = cudiff
				cudata = n.data
			}
			n = n.right
		} else if n.data > val {
			cudiff = n.data - val
			if cudiff < diff {
				diff = cudiff
				cudata = n.data
			}
			n = n.left
		} else {
			return n.data
		}
	}
	return cudata
}

func inorder(root *treeNode) {
	if root == nil {
		return
	}
	inorder(root.left)
	fmt.Print(root.data, "-")
	inorder(root.right)
}
func postorder(root *treeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.data, "-")
	postorder(root.left)
	postorder(root.right)
}
func preorder(root *treeNode) {
	if root == nil {
		return
	}
	preorder(root.left)
	preorder(root.right)
	fmt.Print(root.data, "-")
}
func main() {
	var t Tree
	var arr = []int{2, 1, 3, 5, 6, 10, 24}
	for _, v := range arr {
		t.insert(v)
	}
	inorder(t.root)
	fmt.Println()
	postorder(t.root)
	fmt.Println()
	preorder(t.root)
	fmt.Println("")
	fmt.Println(t.root.Search(2))
	t.delete(30)
	inorder(t.root)
	fmt.Println("")
	levelNodes(t.root, 2)
	fmt.Println("")
	t.leafnodes()
	fmt.Println("")
	t.nonleafnodes()
	fmt.Println("")
	t.rightview()
	fmt.Println("")
	t.leftview()
	fmt.Println("")
	t.closetValue(9)
	t.root = trimBST(t.root, 3, 4)
	inorder(t.root)
}
