package main

import (
	"reflect"
	"testing"
)

func assertPrimeFactor(t *testing.T, n int, expected []int) {
	actual := factorsOf(n)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestPrimeFactors(t *testing.T) {
	t.Run("TestFactors", func(t *testing.T) {
		assertPrimeFactor(t, 1, []int{})
		assertPrimeFactor(t, 2, []int{2})
		assertPrimeFactor(t, 4, []int{2, 2})
		assertPrimeFactor(t, 5, []int{5})
		assertPrimeFactor(t, 6, []int{2, 3})
		assertPrimeFactor(t, 7, []int{7})
		assertPrimeFactor(t, 8, []int{2, 2, 2})
		assertPrimeFactor(t, 9, []int{3, 3})
	})

}
func factorsOf(n int) []int {
	factors := []int{}
	divisor := 2
	for n > 1 {
		for n%divisor == 0 {
			factors = append(factors, divisor)
			n = n / divisor
		}
		divisor++
	}
	return factors
}
