package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Testing slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if want != got {
			t.Errorf("Expected %v but got %v ", want, got)
		}
	})
}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d but got %d ", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %d but got %d ", want, got)
		}
	}
	t.Run("Test for two slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkSums(t, got, want)
	})
	t.Run("Check for an empty Slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 5})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}
