package main

import "fmt"

const arraySize = 7

type hashTable struct {
	array [arraySize]Bucket
}
type Bucket struct {
	head *BucketNode
}
type BucketNode struct {
	key  string
	next *BucketNode
}

//function to get the index of the key, calculating the hash index
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

//inserting into the hash table
func (h *hashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//searching in the hash table
func (h *hashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

//deleting from the hash table
func (h *hashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}
func (b *Bucket) insert(key string) {
	if !b.search(key) {
		var n BucketNode
		n.key = key
		n.next = b.head
		b.head = &n
	} else {
		fmt.Println("Already exsists")
	}
}
func (b *Bucket) search(key string) bool {
	currentnode := b.head
	for currentnode != nil {
		if currentnode.key == key {
			return true
		}
		currentnode = currentnode.next
	}
	return false
}
func (b *Bucket) delete(key string) {
	node := b.head
	if node.key == key {
		b.head = node.next
		node.next = nil
		return
	}
	for node.next != nil {
		if node.next.key == key {
			node.next = node.next.next
			return
		}
		node = node.next
	}
	fmt.Println("Key doesn't exsists")
}
func (h hashTable) display() {
	for v := range h.array {
		ptr := h.array[v].head
		if ptr == nil {
			fmt.Println(ptr)
			continue
		}
		for ptr != nil {
			fmt.Print(ptr.key, "->")
			ptr = ptr.next
		}
		fmt.Println()
	}
}
func main() {
	var h hashTable
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"BUTTERS",
		"RANDY",
	}
	for _, v := range list {
		h.Insert(v)
	}
	fmt.Println(h.Search("RANDY"))
	fmt.Println(h.Search("TITUS"))
	h.Delete("ERIC")
	h.Insert("KYLE")
	h.Insert("JUNMIN")
	fmt.Println(h.Search("RANDY"))
	h.display()
	fmt.Println(h.Search("ERIC"))
	h.Delete("ERIC")
	h.display()
	h.Delete("TITUS")
}
