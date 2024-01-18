package main

import "errors"

var (
	ErrUnderflow = errors.New("underflow")
)

type Stack struct {
	empty    bool
	elements []int
	size     int
}

func NewStack() *Stack {
	return &Stack{
		empty:    true,
		elements: make([]int, 100),
		size:     0,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(i int) {
	s.elements[s.size] = i
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return -1, ErrUnderflow
	}
	s.size--
	return s.elements[s.size], nil
}

func (s *Stack) GetSize() int {
	return s.size
}

func main() {

}
