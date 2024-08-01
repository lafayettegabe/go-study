package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum(numbers): returns the sum of all numbers in the array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("SumAll(numbers): Normal case", func(t *testing.T) {
		numbers := [][]int{{1, 2}, {0, 9}}

		got := SumAll(numbers)
		want := []int{3, 9}

		checkSums(t, got, want)
	})
	t.Run("SumAll(numbers): Empty slice", func(t *testing.T) {
		numbers := [][]int{{}, {0, 9}}

		got := SumAll(numbers)
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("SumAllTails(numbers): Normal case", func(t *testing.T) {
		numbers := [][]int{{1, 2, 3}, {0, 9, 5}}

		got := SumAllTails(numbers)
		want := []int{5, 14}

		checkSums(t, got, want)
	})
	t.Run("SumAllTails(numbers): Empty slice", func(t *testing.T) {
		numbers := [][]int{{}, {0, 9, 5}}

		got := SumAllTails(numbers)
		want := []int{0, 14}

		checkSums(t, got, want)
	})
}

func checkSums(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
