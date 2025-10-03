package main

import "fmt"

type StackInt struct {
	items []int
	size  int
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

func NewStackInt() *StackInt {
	return &StackInt{
		items: make([]int, 0),
	}
}
func (s *StackInt) Add(item int) {
	s.items = append(s.items, item)
	s.size = len(s.items) - 1
}

func (s *StackInt) Pop() int {
	if s.size == 0 {
		return 0
	}
	peek := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.size = len(s.items) - 1
	return peek
}

func (s *StackInt) Size() int {
	return s.size + 1
}

func (s *StackInt) Bottom() int {
	if s.size == 0 {
		return 0
	}
	return s.items[0]

}

func (s *StackInt) Peek() int {
	if s.size == 0 {
		return 0
	}
	return s.items[len(s.items)-1]
}

func (s *StackInt) Print() {
	if s.size == 0 {
		fmt.Println("Empty")
		return
	}
	for i, item := range s.items {
		fmt.Printf("%d: %d\n", i+1, item)
	}
}
