package main

import "fmt"

type item interface{}
type queue struct {
	items []item
}

func (q *queue) enqueue(val item) {
	q.items = append(q.items, val)
}
func (q *queue) dequeue() {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println("Dequeue element is:", q.items[0])
	q.items = q.items[1:]
}
func (q queue) display() {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println(q.items)
}
func main() {
	var q queue
	q.dequeue()
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.display()
	q.dequeue()
	q.enqueue("hello")
	q.display()
	q.dequeue()
	q.display()
}
