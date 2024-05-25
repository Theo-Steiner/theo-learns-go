package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() (perimeter float64) {
	perimeter = (r.Height + r.Width) * 2
	return
}

func (r Rectangle) Area() (area float64) {
	area = r.Height * r.Width
	return
}

func (c Circle) Area() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}
