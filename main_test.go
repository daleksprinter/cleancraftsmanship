package main

import "testing"

func TestCanCreateStack(t *testing.T) {
	s := Stack{}
	actual := s.IsEmpty()
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
