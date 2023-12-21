package main

import "testing"

func TestCanCreateStack(t *testing.T) {
	s := NewStack()
	actual := s.IsEmpty()
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestAfterOnePushIsNotEmpty(t *testing.T) {
	s := NewStack()
	s.Push(0)
	actual := s.IsEmpty()
	expected := false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestAfterOnePushAndOnePopIsNotEmpty(t *testing.T) {
	s := NewStack()
	s.Push(0)
	s.Pop()
	actual := s.IsEmpty()
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
