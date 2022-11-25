package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Iterate("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", i)
	}
}

func ExampleIteration() {
	got := Iterate("b", 1)
	fmt.Println(got)
	// Output b
}
