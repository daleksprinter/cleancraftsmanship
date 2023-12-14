package main

import "testing"

func TestStack(t *testing.T) {
    actual := Stack{}
    expected := Stack{}
    t.Errorf("got %v\nwant %v", actual, expected)
}
