package main

import (
	"testing"
)

var (
	g *Game
)

func setUp() {
	g = NewGame()
}

func rollMany(n int, pins int) {
	for i := 0; i < n; i++ {
		g.Roll(pins)
	}
}

func rollSpare() {
	g.Roll(5)
	g.Roll(5)
}

func rollStrike() {
	g.Roll(10)
}

func TestBowling(t *testing.T) {

	t.Run("TestGutterGame", func(t *testing.T) {
		setUp()
		rollMany(20, 0)
		actual := g.Score()
		expected := 0
		if actual != expected {
			t.Errorf("got %d\nwant %d", actual, expected)
		}
	})

	t.Run("TestAllOnes", func(t *testing.T) {
		setUp()
		rollMany(20, 1)
		actual := g.Score()
		expected := 20
		if actual != expected {
			t.Errorf("got %d\nwant %d", actual, expected)
		}
	})

	t.Run("TestOneSpare", func(t *testing.T) {
		setUp()
		rollSpare()
		g.Roll(7)
		rollMany(17, 0)
		actual := g.Score()
		expected := 24
		if actual != expected {
			t.Errorf("got %d\nwant %d", actual, expected)
		}
	})
	t.Run("TestOneStrike", func(t *testing.T) {
		setUp()
		rollStrike()
		g.Roll(2)
		g.Roll(3)
		rollMany(16, 0)
		actual := g.Score()
		expected := 20
		if actual != expected {
			t.Errorf("got %d\nwant %d", actual, expected)
		}
	})

	t.Run("TestPerfectGame", func(t *testing.T) {
		setUp()
		rollMany(12, 10)
		actual := g.Score()
		expected := 300
		if actual != expected {
			t.Errorf("got %d\nwant %d", actual, expected)
		}
	})

}
