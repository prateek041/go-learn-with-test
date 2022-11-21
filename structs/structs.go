package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct{ // A Rectangle type.
	width float64
	height float64
}

type Circle struct { // A circle type
	radius float64
}

// defining some methods for the types.
func (r Rectangle) Area () float64 { // A method to calculate Area for type Rectangle.
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.height + r.width)
}

func (c Circle) Area () float64{ // A method to calculate Area for type Circle
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}


func Perimeter(rectangle Rectangle) float64{

	height := rectangle.height
	width := rectangle.width

	return 2 * (width + height)
}

func Area(rectangle Rectangle) float64{

	height := rectangle.height
	width := rectangle.width

	return width * height
}