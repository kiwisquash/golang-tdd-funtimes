package array_and_slice

import (
	"reflect"
	"testing"
)

func assertEqual (got, want []int, t testing.TB) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}	
}

func TestSum (t *testing.T) {

	t.Run("add 5 numbers", func (t *testing.T) {
		array := []int{1,2,3,4,5}
		got := SumArray(array)
		want := 15
		if got != want {
			t.Errorf("numbers %v expected %d but got %d", array, want, got)
		}
	})
}

func TestSumAll (t *testing.T) {

	t.Run("multiple slices", func (t *testing.T) {
		slice1 := []int{3,5,7}
		slice2 := []int{2,6,3,-1}

		got := SumAll(slice1, slice2)
		want := []int{15, 10}

		assertEqual(got, want, t)
	})

	t.Run("one slice", func (t *testing.T) {
		slice := []int{3,5,7}

		got := SumAll(slice)
		want := []int{15}

		assertEqual(got, want, t)
	})
}

func TestSumTail (t *testing.T) {

	t.Run("non-emtpy slices", func(t *testing.T) {
		slice1 := []int{3,5,7}
		slice2 := []int{2,6,3,-1}

		got := SumTail(slice1, slice2)
		want := []int{12, 8}

		assertEqual(got, want, t)
	})

	t.Run("use null slice", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{2,6,3,-1}

		got := SumTail(slice1, slice2)
		want := []int{0, 8}

		assertEqual(got, want, t)
	})
}