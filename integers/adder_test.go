package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(4,2)
	fmt.Println(sum)
	// Output: 6
}
func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	if sum != expected {
		t.Errorf("got %d instead of %d", sum, expected)
	}
}