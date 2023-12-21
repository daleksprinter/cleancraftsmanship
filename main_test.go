package main

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	t.Run("TestCanCreateStack", func(t *testing.T) {
		actual := s.IsEmpty()
		expected := true
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestAfterOnePushIsNotEmpty", func(t *testing.T) {
		s.Push(0)
		actual := s.IsEmpty()
		expected := false
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestAfterOnePushAndOnePopIsNotEmpty", func(t *testing.T) {
		s.Push(0)
		s.Pop()
		actual := s.IsEmpty()
		expected := true
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})
}
