package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 3 numbers", func(t *testing.T) {
		numbers := [3]int{1, 2, 3}
		got := Sum(numbers[:])
		expected := 6
		if expected != got {
			t.Errorf("Expected %d but got %d", expected, got)
		}

	})
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6
		if expected != got {
			t.Errorf("Expected %d but got %d", expected, got)
		}
	})
}
func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}
	got := SumAll(numbers1, numbers2)
	expected := []int{6, 15}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}
func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}

	t.Run("sums of some slices", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{4, 5, 6}
		got := SumAllTails(numbers1, numbers2)
		expected := []int{5, 11}
		checkSums(t, got, expected)
	})

	t.Run("sum safely empty slices", func(t *testing.T) {
		numbers1 := []int{}
		numbers2 := []int{4, 5, 6}
		got := SumAllTails(numbers1, numbers2)
		expected := []int{0, 11}
		checkSums(t, got, expected)
	})
}
