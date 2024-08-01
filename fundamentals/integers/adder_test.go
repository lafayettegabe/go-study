package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	expected := 4
	sum := Add(2, 2)

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
