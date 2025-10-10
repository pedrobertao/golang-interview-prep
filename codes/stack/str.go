package main

import "fmt"

// StackString implements a stack data structure for strings using a slice.
// A stack follows Last-In-First-Out (LIFO) principle.
type StackString struct {
	items []string // internal slice to store stack items
}

func main() {
	s := NewStackString()
	s.Add("Hello")
	s.Add("World")
	s.Add("From golang")
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

// NewStackString creates and returns a new empty StackString.
func NewStackString() *StackString {
	return &StackString{
		items: make([]string, 0),
	}
}

// Add pushes a new item onto the top of the stack.
// Time complexity: O(1) amortized
func (s *StackString) Add(item string) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// Returns empty string if the stack is empty.
// Time complexity: O(1)
func (s *StackString) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	peek := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return peek
}

// Size returns the number of items in the stack.
// Time complexity: O(1)
func (s *StackString) Size() int {
	return len(s.items)
}

// Bottom returns the bottom (first) item in the stack without removing it.
// Returns empty string if the stack is empty.
// Time complexity: O(1)
func (s *StackString) Bottom() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[0]
}

// Peek returns the top item in the stack without removing it.
// Returns empty string if the stack is empty.
// Time complexity: O(1)
func (s *StackString) Peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

// Print displays all items in the stack from bottom to top.
// Prints "Empty" if the stack is empty.
// Time complexity: O(n) where n is the number of items
func (s *StackString) Print() {
	if len(s.items) == 0 {
		fmt.Println("Empty")
		return
	}
	for i, item := range s.items {
		fmt.Printf("%d: %s\n", i+1, item)
	}
}
