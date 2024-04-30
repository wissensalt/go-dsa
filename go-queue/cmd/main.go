package main

import (
	"fmt"
)

type Queue struct {
	Items []interface{}
}

type QueueImplementation interface {
	Enqueue(item interface{})
	Dequeue() interface{}
	Head() interface{}
	Tail() interface{}
	IsEmpty() bool
	Size() int
	Print()
}

// Enqueue Adds an element to the back of the queue.
func (q *Queue) Enqueue(item interface{}) {
	q.Items = append(q.Items, item)
}

// Dequeue Removes and returns the front element from the queue.
func (q *Queue) Dequeue() interface{} {
	if len(q.Items) == 0 {
		return nil
	}

	itemToRemove := q.Items[0]
	q.Items = q.Items[1:]

	return itemToRemove
}

// Head Returns the front element from the queue without removing it
func (q *Queue) Head() interface{} {
	if len(q.Items) == 0 {
		return nil
	}

	return q.Items[0]
}

// Tail Returns the rear element from the queue without removing it
func (q *Queue) Tail() interface{} {
	if len(q.Items) == 0 {
		return nil
	}

	return q.Items[len(q.Items)-1]
}

// IsEmpty Checks if the queue is empty.
func (q Queue) IsEmpty() bool {

	return len(q.Items) == 0
}

// Size Returns the number of elements in the queue.
func (q Queue) Size() int {

	return len(q.Items)
}

func (q *Queue) Print() {
	fmt.Printf("\n===QUEUE===\n")
	for _, v := range q.Items {
		fmt.Printf("%v ", v)
	}
}

func main() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Print()
	queue.Dequeue()
	queue.Print()
	fmt.Println("Size: ", queue.Size())
	fmt.Println("Is Empty: ", queue.IsEmpty())
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println("Is Empty: ", queue.IsEmpty())
}
