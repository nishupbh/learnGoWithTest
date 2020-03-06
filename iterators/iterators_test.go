package iterators

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	got := iterator("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

// ExampleRepeat for testing using example
func ExampleRepeat() {
	val := iterator("b", 5)
	fmt.Println(val)
	// Output: bbbbb
}
