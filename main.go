package main

type Stack struct {
	empty bool
	size  int
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
}

func (s *Stack) Pop() int {
	s.size--
	return -1
}

func (s *Stack) GetSize() int {
	return s.size
}

func main() {

}
