package structs

import "math"

//Rectangle struct for the shape of rectange
type Rectangle struct {
	height float64
	width  float64
}

//Circle struct for the radius of circle
type Circle struct {
	radius float64
}

//Triangle struct for the radius of circle
type Triangle struct {
	height float64
	base   float64
}

// Shape interface for passing the shape in area method
type Shape interface {
	area() float64
}

func perimeter(r Rectangle) float64 {

	return 2 * (r.height + r.width)
}

func (r Rectangle) area() float64 {
	return r.height * r.width

}

func (r Circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

func (t Triangle) area() float64 {
	return 0.5 * t.height * t.base
}
