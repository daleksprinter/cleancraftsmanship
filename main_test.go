package main

import "testing"

func TestStack(t *testing.T) {
	t.Run("TestCanCreateStack", func(t *testing.T) {
		s := NewStack()
		actual := s.IsEmpty()
		expected := true
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestAfterOnePushIsNotEmpty", func(t *testing.T) {
		s := NewStack()
		s.Push(0)
		actual := s.IsEmpty()
		expected := false
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}

		actual2 := s.GetSize()
		expected2 := 1
		if actual2 != expected2 {
			t.Errorf("got %v\nwant %v", actual2, expected2)
		}
	})

	t.Run("TestAfterOnePushAndOnePopIsEmpty", func(t *testing.T) {
		s := NewStack()
		s.Push(0)
		s.Pop()
		actual := s.IsEmpty()
		expected := true
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}

		actual2 := s.GetSize()
		expected2 := 0
		if actual2 != expected2 {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TesstAfterTwoPushesSizeIsTwo", func(t *testing.T) {
		s := NewStack()
		s.Push(0)
		s.Push(0)
		actual := s.GetSize()
		expected := 2
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("poppingEmptyStack_returnsErrUnderflow", func(t *testing.T) {
		s := NewStack()
		_, actual := s.Pop()
		expected := ErrUnderflow
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("afterPushingX_willPopX", func(t *testing.T) {
		s := NewStack()
		s.Push(99)
		actual, _ := s.Pop()
		expected := 99
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}

		s.Push(88)
		actual2, _ := s.Pop()
		expected2 := 88
		if actual2 != expected2 {
			t.Errorf("got %v\nwant %v", actual2, expected2)
		}
	})
}
