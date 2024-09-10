package internal

import "fmt"

type Stack struct {
	Elements []int
	Capacity int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		Elements: make([]int, 0),
		Capacity: capacity,
	}
}

func (s *Stack) Push(element int) {
	if len(s.Elements) == s.Capacity {
		fmt.Println("Stack overflow")
		return
	}
	s.Elements = append(s.Elements, element)
}

func (s *Stack) Pop() int {
	if len(s.Elements) == 0 {
		fmt.Println("Stack underflow")
		return 0
	}
	element := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return element
}

func (s *Stack) Top() int {
	if len(s.Elements) == 0 {
		fmt.Println("Stack is empty")
		return 0
	}
	return s.Elements[len(s.Elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func (s *Stack) Size() int {
	return len(s.Elements)
}

func (s *Stack) Print() {
	for i := len(s.Elements) - 1; i >= 0; i-- {
		fmt.Println(s.Elements[i])
	}
}
