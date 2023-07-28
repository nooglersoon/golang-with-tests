package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// It is a convention in Go to have the receiver variable be the first letter of the type.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}

type Circle struct {
	radius float64
}

// It is a convention in Go to have the receiver variable be the first letter of the type.
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
