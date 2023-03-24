package main

import "fmt"

const arraySize = 26

//node of a trie

type node struct {
	children [arraySize]*node
	isEnd    bool
}
type Trie struct {
	root *node
}

//inserting into the trie

func (t *Trie) insert(word string) {
	wordlength := len(word)
	ptr := t.root
	for i := 0; i < wordlength; i++ {
		charIndex := word[i] - 'a'
		if ptr.children[charIndex] == nil {
			ptr.children[charIndex] = &node{}
		}
		ptr = ptr.children[charIndex]
	}
	ptr.isEnd = true
}

//deleting from a trie

func (t *Trie) delete(word string) {
	t.root.deleteHelper(word, 0)
}
func (n *node) deleteHelper(word string, depth int) *node {
	if n == nil {
		return nil
	}
	if depth == len(word) {
		if n.isEnd {
			n.isEnd = false
		}
		if !n.check() {
			n = nil
		}
		return n
	}
	charIndex := word[depth] - 'a'
	if n.children[charIndex] == nil {
		fmt.Println("Word doesn't exists")
		return n
	}
	n.children[charIndex] = n.children[charIndex].deleteHelper(word, depth+1)
	if !n.check() && n.isEnd == false {
		n = nil
	}
	return n
}

func (n *node) check() bool {
	for i := 0; i < arraySize; i++ {
		if n.children[i] != nil {
			return true
		}
	}
	return false
}

//searching in a trie

func (t *Trie) search(word string) bool {
	wordlength := len(word)
	ptr := t.root
	for i := 0; i < wordlength; i++ {
		charIndex := word[i] - 'a'
		if ptr.children[charIndex] == nil {
			return false
		}
		ptr = ptr.children[charIndex]
	}
	return ptr.isEnd
}


func main() {
	var t Trie
	t.root = &node{}
	words := []string{
		"rain",
		"ran",
		"raining",
		"or",
		"oreo",
		"organ",
		"stem",
		"steak",
		"stable",
	}
	for _, v := range words {
		t.insert(v)
	}
	fmt.Println(t.search("stable"))
	t.delete("raining")
	fmt.Println(t.search("raining"))
}
