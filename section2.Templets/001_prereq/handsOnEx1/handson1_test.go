package main

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("area of a square", func(t *testing.T) {
		s := square{Height: 5.5}

		got := s.Area()
		want := 30.25

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("area of a circle", func(t *testing.T) {
		c := circle{Radius: 3.2}

		got := c.Area()
		want := 32.1699

		if !closeEnough(got, want) {
			t.Errorf("got %.4f want %.4f", got, want)
		}
	})
} //FAIL: got 32.1699 want 32.1699 ??

// closeEnough checks if two floating-point numbers are close enough within a small margin of error.
func closeEnough(a, b float64) bool {
	epsilon := 0.0001              // Define a small margin of error
	return math.Abs(a-b) < epsilon // ha 2 float abszolut ertekenek a kulonbsege kisebb, mint 0.0001 akkor true
}

func TestInfo(t *testing.T) {
	s := square{Height: 4.0}

	got := Info(s)
	want := 16.0

	if !closeEnough(got, want) {
		t.Errorf("got %.4f want %.4f", got, want)
	}
}
