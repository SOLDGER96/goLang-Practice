package main

import (
	"fmt"
)

type Queue struct{
	items []string
}

func (q *Queue) Enqueue(item string) {
	q.items = append(q.items, item)
	fmt.Println("Queue: ", *q)
}

func (q *Queue) Dequeue() string {
	// If queue is empty
	if len(q.items) == 0 {
		fmt.Println("Queue is Empty!")
		return ""
	}
	topItem := q.items[0]
	
	q.items = q.items[1:]
	fmt.Println("Queue: ", *q)
	fmt.Println("Item Poped: ", topItem)
	return topItem
}

func main() {
	myQueue := Queue{}
	myQueue.Dequeue()
	myQueue.Enqueue("Abc")
	myQueue.Enqueue("tanm")
	myQueue.Enqueue("jod")
	myQueue.Dequeue()
	myQueue.Enqueue("raj")
}