package main

import (
	"fmt"
)

type Node struct {
	Item interface{}
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

type DoublyLinkedListImplementation interface {
	Append(item interface{})
	PrintForward()
	PrintBackward()
}

func (dll *DoublyLinkedList) Append(item interface{}) {
	newNode := &Node{item, nil, nil}
	if dll.Head == nil && dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode

		return
	}

	dll.Tail.Next = newNode
	newNode.Prev = dll.Tail
	dll.Tail = newNode
}

func (dll *DoublyLinkedList) PrintForward() {
	fmt.Println("Print Forward")
	current := dll.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Item)
		current = current.Next
	}
	fmt.Println()
}

func (dll *DoublyLinkedList) PrintBackward() {
	fmt.Println("Print Backward")
	current := dll.Tail
	for current != nil {
		fmt.Printf("%v -> ", current.Item)
		current = current.Prev
	}
	fmt.Println()
}

func main() {
	dll := &DoublyLinkedList{}
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Append(4)
	dll.PrintForward()
	dll.PrintBackward()
}
