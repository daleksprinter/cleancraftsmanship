package main

import "errors"

var (
	ErrUnderflow = errors.New("underflow")
)

type Stack struct {
	empty   bool
	element int
	size    int
}

func NewStack() *Stack {
	return &Stack{
		empty: true,
		size:  0,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(i int) {
	s.size++
	s.element = i
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return -1, ErrUnderflow
	}
	s.size--
	return s.element, nil
}

func (s *Stack) GetSize() int {
	return s.size
}

func main() {

}
