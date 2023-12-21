package main

type Stack struct {
	empty bool
}

func NewStack() *Stack {
	return &Stack{empty: true}
}

func (s *Stack) IsEmpty() bool {
	return s.empty
}

func (s *Stack) Push(i int) {
	s.empty = false
}

func main() {

}
