package structs

import "math"

// Rectangle ... Struct fro rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle ...
type Circle struct {
	Radius float64
}

// Shape ...
type Shape interface {
	Area() float64
}

// Area ...
func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

// Area ...
func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}

// Perimeter ...
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
