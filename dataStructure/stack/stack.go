package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
	fmt.Println("Stack: ", *s)
}

func (s *Stack) Pop() string {
	lastIndex := len(s.items) - 1
	if lastIndex == -1 {
		fmt.Println("No items in Stack!")
		return ""
	}
	topItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	fmt.Println("Stack: ", *s)
	fmt.Println("Item Poped: ", topItem)
	return topItem
}

func main() {
	myStack := Stack{}

	myStack.Push("Tanmay")
	myStack.Push("jod")
	myStack.Pop()
}


