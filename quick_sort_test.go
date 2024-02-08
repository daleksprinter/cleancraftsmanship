package main

import (
	"reflect"
	"testing"
)

func quick_sort(list []int) []int {
	if len(list) == 0 {
		return list
	}
	middle := list[0]
	lessers := []int{}
	greaters := []int{}
	for _, v := range list {
		if v < middle {
			lessers = append(lessers, v)
		}
	}
	for _, v := range list {
		if v > middle {
			greaters = append(greaters, v)
		}
	}
	middles := []int{}
	for _, v := range list {
		if v == middle {
			middles = append(middles, v)
		}
	}
	result := []int{}
	result = append(result, quick_sort(lessers)...)
	result = append(result, middles...)
	result = append(result, quick_sort(greaters)...)

	return result
}

func TestQuickSort(t *testing.T) {
	t.Run("TestSort", func(t *testing.T) {
		actual := quick_sort([]int{1})
		expected := []int{1}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := quick_sort([]int{1, 2})
		expected := []int{1, 2}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := quick_sort([]int{2, 1})
		expected := []int{1, 2}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := quick_sort([]int{2, 1, 3})
		expected := []int{1, 2, 3}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := quick_sort([]int{2, 3, 1})
		expected := []int{1, 2, 3}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := quick_sort([]int{5, 3, 1, 2, 4})
		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := quick_sort([]int{1, 3, 1, 2})
		expected := []int{1, 1, 2, 3}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})

	t.Run("TestSort", func(t *testing.T) {
		actual := quick_sort([]int{1, 3, 1, 2, 4})
		expected := []int{1, 1, 2, 3, 4}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	})
}
