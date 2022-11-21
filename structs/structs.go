package structs

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