package array_and_slice

import "testing"

func TestArraySum (t *testing.T) {

	t.Run("add 5 numbers", func (t *testing.T) {
		array := []int{1,2,3,4,5}
		got := SumArray(array)
		want := 15
		if got != want {
			t.Errorf("numbers %v expected %d but got %d", array, want, got)
		}
	})
}