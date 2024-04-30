package main

import (
	"fmt"
)

type Stack struct {
	Items []interface{}
}

type StackImplementation interface {
	Push(item interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
	Size() int
	Print()
}

// Push Adding an element to the top element of stack
func (s *Stack) Push(value interface{}) {
	s.Items = append(s.Items, value)
}

// Pop Remove and Return the top element of stack
func (s *Stack) Pop() interface{} {
	if len(s.Items) == 0 {
		return nil
	}

	itemForRemoval := s.Items[len(s.Items)-1]
	// Removing the item
	s.Items = s.Items[:len(s.Items)-1]

	return itemForRemoval
}

// Peek Returning top element of stack
func (s *Stack) Peek() interface{} {
	if len(s.Items) == 0 {
		return nil
	}

	return s.Items[len(s.Items)-1]
}

// IsEmpty Check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

// Size Return size of stack
func (s *Stack) Size() int {
	return len(s.Items)
}

func (s *Stack) Print() {
	for i := len(s.Items) - 1; i >= 0; i-- {
		fmt.Println(s.Items[i])
	}
}

func main() {
	stack := &Stack{nil}
	fmt.Println("Check if stack empty: ", stack.IsEmpty())
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Print()
	currentTop := stack.Peek()
	fmt.Println("Top Element:", currentTop)
	removedItem := stack.Pop()
	fmt.Println("Removed Element: ", removedItem)
	stack.Print()
	currentSize := stack.Size()
	fmt.Println("Current Size:", currentSize)
	fmt.Println("Check if stack empty: ", stack.IsEmpty())
}
