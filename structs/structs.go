package structs

import "math"

type Shape interface {
	Area() float64
}

// All about triangle
type Triangle struct { // A type triangle
	base float64
	height float64
}

func (t Triangle) Area() float64{
	return 0.5 * t.base * t.height
}

// All about rectangles.
type Rectangle struct{ // A Rectangle type.
	width float64
	height float64
}

func (r Rectangle) Area () float64 { // A method to calculate Area for type Rectangle.
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64{ // A method to calculate Perimeter for type Rectangle
	return 2*(r.height + r.width)
}

// All about Circle.
type Circle struct { // A circle type
	radius float64
}

func (c Circle) Area () float64{ // A method to calculate Area for type Circle
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 { // A method to calculate Perimeter for type Circle
	return 2 * math.Pi * c.radius
}
