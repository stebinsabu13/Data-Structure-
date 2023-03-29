package main

import "fmt"

type node struct {
	data interface{}
	next *node
}
type queue struct {
	front *node
	back  *node
}

func (q *queue) enqueue(val interface{}) {
	var n node
	n.data = val
	if q.back == nil {
		q.front = &n
		q.back = &n
		return
	}
	q.back.next = &n
	q.back = &n
}
func (q *queue) dequeue() {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return
	}
	q.front = q.front.next
}
func (q queue) display() {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return
	}
	ptr := q.front
	for ptr != nil {
		fmt.Print(ptr.data, " ")
		ptr = ptr.next
	}
	fmt.Println()
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
