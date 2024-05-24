package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	got := Sum(1, 2, 3, 4, 5)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given", got, want)
	}
}

func TestArraySum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := ArraySum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSliceSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}

	got := SliceSum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{3, 6})
	want := []int{3, 9}

	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 6})
		want := []int{2, 6}

		checkSums(t, got, want)
	})

	t.Run("safely handle empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 6})
		want := []int{0, 6}

		checkSums(t, got, want)
	})
}
