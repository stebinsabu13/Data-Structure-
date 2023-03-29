package main

import "fmt"

type item interface{}

type stack struct {
	items []item
}

func (s *stack) push(val item) {
	s.items = append(s.items, val)
}
func (s *stack) pop() {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Popped element", s.items[len(s.items)-1])
	s.items = s.items[:len(s.items)-1]
}
func (s stack) display() {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println(s.items)
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
