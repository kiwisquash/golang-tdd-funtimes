package main

import "math"

type Rectangle struct {
	Length float64
	Width float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
} 

func (r Rectangle) Perimeter() float64 {
	return 2*r.Length + 2*r.Width
}

func (c Circle) Perimeter() float64 {
	return 2*c.Radius*math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Length*r.Width
}

func (c Circle) Area() float64 {
	return math.Pi*math.Pow(c.Radius, 2)
}

