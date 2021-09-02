package structs

import "math"

type Rectangle struct {
	Height float64
	Width  float64
}

func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.Width + rec.Height)
}

func (rec Rectangle) Area() float64 {
	return rec.Height * rec.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

var _ Shape = Rectangle{}
var _ Shape = Circle{}
