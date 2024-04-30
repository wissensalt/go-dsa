package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type NodeImplementation interface {
	insert(val int)
	searchLeft(val int) bool
	searchRight(val int) bool
	isComplete() bool
	switchPosition()
}

func (n *Node) insert(val int) {
	if n.Value == 0 {
		n.Value = val

		return
	}

	if !n.isComplete() {
		if n.Left == nil {
			n.Left = &Node{val, nil, nil}

			return
		}

		if n.Right == nil {
			n.Right = &Node{val, nil, nil}

			if n.Left.Value > n.Right.Value {
				n.switchPosition()
			}

			return
		}
	} else if !n.Left.isComplete() {
		n.Left.insert(val)
	} else if !n.Right.isComplete() {
		n.Right.insert(val)
	} else {
		n.Left.insert(val)
	}
}

func (n *Node) isComplete() bool {

	return n.Left != nil && n.Left.Value != 0 && n.Right != nil && n.Right.Value != 0
}

func (n *Node) switchPosition() {
	tempValue := n.Left.Value
	n.Left.Value = n.Right.Value
	n.Right.Value = tempValue
}

func (n *Node) searchLeft(val int) bool {
	if n == nil {
		return false
	}

	if n.Value == val {
		return true
	}

	if n.Left != nil {
		return n.Left.searchLeft(val)
	}

	return false
}

func (n *Node) searchRight(val int) bool {
	if n == nil {
		return false
	}

	if n.Value == val {
		return true
	}

	if n.Right != nil {
		return n.Right.searchLeft(val)
	}

	return false
}

func (n *Node) search(val int) bool {
	if n.searchLeft(val) {
		return true
	}

	return n.searchRight(val)
}

func main() {
	root := &Node{}

	root.insert(1)
	root.insert(2)
	root.insert(3)
	root.insert(4)
	root.insert(5)
	root.insert(6)
	root.insert(7)
	//root.insert(11)
	//root.insert(12)
	//root.insert(13)
	//root.insert(14)
	//root.insert(15)
	//root.insert(16)

	fmt.Println(root.search(2))
	fmt.Println(root.search(10))
	fmt.Println(root.search(3))
	fmt.Println(root.search(4))
	//fmt.Println(root)
}
