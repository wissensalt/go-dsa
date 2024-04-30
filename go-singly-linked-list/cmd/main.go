package main

import (
	"fmt"
)

type Node struct {
	Item interface{}
	Next *Node
}

type SinglyLinkedListImplementation interface {
	Append(item interface{})
	Print()
	Search(item interface{}) bool
}

func (n *Node) Append(item interface{}) {
	if n == nil {
		n = &Node{item, nil}

		return
	}

	if n.Item == nil {
		n.Item = item

		return
	}

	for {
		if n.Next == nil {
			n.Next = &Node{item, nil}

			break
		}

		n = n.Next
	}
}

func (n *Node) Print() {
	for {
		fmt.Printf("%v -> ", n.Item)
		if n.Next == nil {
			fmt.Println()
			break
		}

		n = n.Next
	}
}

func (n *Node) Search(item interface{}) bool {
	for {
		if n.Item == item {
			return true
		}

		if n.Next == nil {
			return false
		}

		n = n.Next
	}
}

func main() {
	node := &Node{}
	node.Append(1)
	node.Append(2)
	node.Append(3)
	node.Append(4)
	node.Append(5)
	node.Print()
	isVal3Found := node.Search(3)
	fmt.Println("Value 3 Found ?", isVal3Found)

	isVal6Found := node.Search(6)
	fmt.Println("Value 6 Found ?", isVal6Found)
}
