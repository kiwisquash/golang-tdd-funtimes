package main

import (
	"math"
	"testing"
)

func TestShape(t *testing.T) {

	tests := []struct {
		shape Shape
		areaWant float64
		perimeterWant float64
	}{
		{Rectangle{3, 4}, 12.0, 1.40},
		{Circle{1}, math.Pi, 2*math.Pi},
	}

	for _, tt := range tests {
		gotArea := tt.shape.Area()
		if gotArea != tt.areaWant {
			t.Errorf("For area, got %g want %g", gotArea, tt.areaWant)
		}

		gotPerimeter := tt.shape.Perimeter()
		if gotPerimeter != tt.perimeterWant {
			t.Errorf("For perimeter, got %g want %g", gotPerimeter, tt.perimeterWant)
		}

	}
}