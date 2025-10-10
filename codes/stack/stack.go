package main

import "fmt"

// ItemType defines the constraint for types that can be used in GenericStack.
// Supports all numeric types and strings.
type ItemType interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		int | ~int8 | ~int16 | ~int32 | ~int64 |
		float32 | ~float64 | ~string
}

// GenericStack implements a type-safe generic stack data structure using Go generics.
// A stack follows Last-In-First-Out (LIFO) principle.
type GenericStack[T ItemType] struct {
	items []T // internal slice to store stack items
	zero  T   // zero value of type T, used for returning defaults on empty stack
}

// NewGenericStack creates and returns a new empty GenericStack of type T.
func NewGenericStack[T ItemType]() *GenericStack[T] {
	return &GenericStack[T]{
		items: make([]T, 0),
	}
}

// Add pushes a new item onto the top of the stack.
// Time complexity: O(1) amortized
func (s *GenericStack[T]) Add(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// Returns zero value of type T if the stack is empty.
// Time complexity: O(1)
func (s *GenericStack[T]) Pop() T {
	if len(s.items) == 0 {
		return s.zero
	}
	peek := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return peek
}

// Size returns the number of items in the stack.
// Time complexity: O(1)
func (s *GenericStack[T]) Size() int {
	return len(s.items)
}

// Bottom returns the bottom (first) item in the stack without removing it.
// Returns zero value of type T if the stack is empty.
// Time complexity: O(1)
func (s *GenericStack[T]) Bottom() T {
	if len(s.items) == 0 {
		return s.zero
	}
	return s.items[0]
}

// Peek returns the top item in the stack without removing it.
// Returns zero value of type T if the stack is empty.
// Time complexity: O(1)
func (s *GenericStack[T]) Peek() T {
	if len(s.items) == 0 {
		return s.zero
	}
	return s.items[len(s.items)-1]
}

// Print displays all items in the stack from bottom to top.
// Prints "Empty" if the stack is empty.
// Time complexity: O(n) where n is the number of items
func (s *GenericStack[T]) Print() {
	if len(s.items) == 0 {
		fmt.Println("Empty")
		return
	}
	for i, item := range s.items {
		fmt.Printf("%d: %v\n", i+1, item)
	}
}

func main() {
	fmt.Println("String ================")
	ExampleString()
	fmt.Println("Int ===================")
	ExampleInt()
}

// ExampleString demonstrates the GenericStack usage with string type.
func ExampleString() {
	s := NewGenericStack[string]()
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

// ExampleInt demonstrates the GenericStack usage with int type.
func ExampleInt() {
	s := NewGenericStack[int]()
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
