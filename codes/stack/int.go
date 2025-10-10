package main

import "fmt"

// StackInt implements a stack data structure for integers using a slice.
// A stack follows Last-In-First-Out (LIFO) principle.
type StackInt struct {
	items []int // internal slice to store stack items
}

func main() {
	s := NewStackInt()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Print()
	fmt.Println("Peek:", s.Peek())
	fmt.Println("Bottom:", s.Bottom())
	fmt.Println("Size:", s.Size())
	popped := s.Pop()
	fmt.Println("Popped:", popped)
	fmt.Println("Peek:", s.Peek())
	fmt.Println("Bottom:", s.Bottom())
	fmt.Println("Size:", s.Size())
	s.Pop()
	s.Pop()
	s.Print()
	fmt.Println("Peek:", s.Peek())
	fmt.Println("Bottom:", s.Bottom())
	fmt.Println("Size:", s.Size())
}

// NewStackInt creates and returns a new empty StackInt.
func NewStackInt() *StackInt {
	return &StackInt{
		items: make([]int, 0),
	}
}
// Add pushes a new item onto the top of the stack.
// Time complexity: O(1) amortized
func (s *StackInt) Add(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// Returns -1 if the stack is empty.
// Time complexity: O(1)
func (s *StackInt) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	peek := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return peek
}

// Size returns the number of items in the stack.
// Time complexity: O(1)
func (s *StackInt) Size() int {
	return len(s.items)
}

// Bottom returns the bottom (first) item in the stack without removing it.
// Returns 0 if the stack is empty.
// Time complexity: O(1)
func (s *StackInt) Bottom() int {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[0]
}

// Peek returns the top item in the stack without removing it.
// Returns 0 if the stack is empty.
// Time complexity: O(1)
func (s *StackInt) Peek() int {
	if len(s.items) == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

// Print displays all items in the stack from bottom to top.
// Prints "Empty" if the stack is empty.
// Time complexity: O(n) where n is the number of items
func (s *StackInt) Print() {
	if len(s.items) == 0 {
		fmt.Println("Empty")
		return
	}
	for i, item := range s.items {
		fmt.Printf("%d: %d\n", i+1, item)
	}
}
