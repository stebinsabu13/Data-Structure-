package main

import "fmt"

type node struct {
	data interface{}
	next *node
}
type stack struct {
	top    *node
	length int
}

func (s *stack) push(val interface{}) {
	var n node
	n.data = val
	if s.length == 0 {
		s.top = &n
		s.length++
		return
	}
	n.next = s.top
	s.top = &n
	s.length++
}
func (s *stack) pop() {
	if s.length == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Popped element is:", s.top.data)
	s.top = s.top.next
	s.length--
}
func (s stack) display() {
	if s.length == 0 {
		fmt.Println("Stack is empty")
		return
	}
	ptr := s.top
	for i := 0; i < s.length; i++ {
		fmt.Print(ptr.data, " ")
		ptr = ptr.next
	}
	fmt.Println()
}
func main() {
	var s stack
	s.pop()
	s.push(5)
	s.push(10)
	s.push(15)
	s.display()
	s.pop()
	s.display()
	s.push(20)
	s.display()
	s.push("hello")
	s.display()
	s.pop()
	s.display()
}
