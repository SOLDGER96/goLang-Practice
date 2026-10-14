# Data Structures in Go: Stack and Queue

This document contains the custom implementations of a Stack (LIFO) and a Queue (FIFO) using Go slices.

## 1. Stack (LIFO - Last-In, First-Out)
A Stack works like a stack of plates. The last item added is the first one removed.

```go
package main

import "fmt"

type Stack struct {
    items []string
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item string) {
    s.items = append(s.items, item)
    fmt.Println("Stack: ", *s)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() string {
    if len(s.items) == 0 {
        fmt.Println("Stack is Empty!")
        return ""
    }
    
    lastIndex := len(s.items) - 1
    topItem := s.items[lastIndex]
    s.items = s.items[:lastIndex]
    
    fmt.Println("Stack: ", *s)
    fmt.Println("Item Popped: ", topItem)
    return topItem
}
```

## 2. Queue (FIFO - First-In, First-Out)
A Queue works like a line at a movie theater. The first person to join the line is the first one to leave.

```go
type Queue struct {
    items []string
}

// Enqueue adds an item to the back of the line
func (q *Queue) Enqueue(item string) {
    q.items = append(q.items, item)
    fmt.Println("Queue: ", *q)
}

// Dequeue removes and returns the front item from the line
func (q *Queue) Dequeue() string {
    if len(q.items) == 0 {
        fmt.Println("Queue is Empty!")
        return ""
    }
    
    frontItem := q.items[0]
    q.items = q.items[1:]
    
    fmt.Println("Queue: ", *q)
    fmt.Println("Item Dequeued: ", frontItem)
    return frontItem
}
```