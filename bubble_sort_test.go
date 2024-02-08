package main

import (
	"reflect"
	"testing"
)

func bubble_sort(list []int) []int {
	if len(list) > 1 {
		for limit := len(list) - 1; limit > 0; limit-- {
			for i := 0; i < limit; i++ {
				j := i + 1
				if list[i] > list[j] {
					list[i], list[j] = list[j], list[i]
				}
			}
		}

	}
	return list
}

func TestBubbleSort(t *testing.T) {
	t.Run("TestSort", func(t *testing.T) {
		actual := bubble_sort([]int{1})
		expected := []int{1}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := bubble_sort([]int{1, 2})
		expected := []int{1, 2}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := bubble_sort([]int{2, 1})
		expected := []int{1, 2}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := bubble_sort([]int{2, 1, 3})
		expected := []int{1, 2, 3}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := bubble_sort([]int{2, 3, 1})
		expected := []int{1, 2, 3}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := bubble_sort([]int{5, 3, 1, 2, 4})
		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})
}
