package iteration

import (
	"fmt"
	"testing"
)

func assertEqual(repeated, expected string, t testing.TB) {
	if repeated != expected {
		t.Errorf("got %q when %q was expected", repeated, expected)
	}
}

func ExampleRepeat() {
	repeated := Repeat("example", 2)
	fmt.Println(repeated)
	// Output: exampleexample
}

func TestRepeat(t *testing.T) {

	t.Run("repeat multiple times", func (t *testing.T) {
		repeated := Repeat("ki", 6)
		expected := "kikikikikiki"
		assertEqual(repeated, expected, t)
	})	

	t.Run("edge case ", func (t *testing.T) {
		repeated := Repeat("ki", 0)
		expected := ""
		assertEqual(repeated, expected, t)

	})
}

func BenchmarkRepeat(b *testing.B) {
	i := 0
	for i < b.N {
		i++
		Repeat("a", 6)
	}
}