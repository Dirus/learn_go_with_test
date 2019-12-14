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
