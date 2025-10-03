package main

import "fmt"

type StatckString struct {
	items []string
	size  int
}

func main() {
	s := NewStackString()
	s.Add("Hhello")
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

func NewStackString() *StatckString {
	return &StatckString{
		items: make([]string, 0),
	}
}
func (s *StatckString) Add(item string) {
	s.items = append(s.items, item)
	s.size = len(s.items) - 1
}

func (s *StatckString) Pop() string {
	if s.size == 0 {
		return ""
	}
	peek := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.size = len(s.items) - 1
	return peek
}

func (s *StatckString) Size() int {
	return s.size
}

func (s *StatckString) Bottom() string {
	if s.size == 0 {
		return ""
	}
	return s.items[0]

}

func (s *StatckString) Peek() string {
	if s.size == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

func (s *StatckString) Print() {
	if s.size == 0 {
		fmt.Println("Empty")
	}
	for i, r := range s.items {
		fmt.Printf("%d: %s\n", i+1, string(r))
	}
}
