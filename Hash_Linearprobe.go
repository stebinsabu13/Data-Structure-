package main

import "fmt"

const arraySize = 7

type hashTable struct {
	arr [arraySize]int
}

func hash(val int) int {
	return val % arraySize
}
func (h *hashTable) Insert(val int) {
	index := hash(val)
	if h.arr[index] == 0 {
		h.arr[index] = val
	} else {
		for i := index + 1; i < arraySize; i++ {
			if h.arr[i] == 0 {
				h.arr[i] = val
				break
			}
		}
	}
}
func main() {
	h := hashTable{}
	var array = []int{50, 700, 76, 85, 92, 73, 101}
	for _, v := range array {
		h.Insert(v)
	}
	fmt.Println(h.arr)
}
