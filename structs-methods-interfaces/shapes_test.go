package main

import "testing"



func TestPerimeter(t *testing.T) {

	checkPermieter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("perimeter of rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		want := 60.0

		checkPermieter(t, rectangle, want)

	})


	t.Run("perimeter (aka circumference) of circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 62.83185307179586

		checkPermieter(t, circle, want)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("area of rectangle", func(t *testing.T) {

		rectangle := Rectangle{10.0, 20.0}
		want := 200.0

		checkArea(t, rectangle, want)

	}) 


	t.Run("area of circle", func(t *testing.T) {

		circle := Circle{10.0}
		want := 314.1592653589793

		checkArea(t, circle, want)
	}) 
	
}