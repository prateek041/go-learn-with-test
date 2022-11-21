package structs

import "testing"

func TestPerimeter(t *testing.T){
 	
	rectangle:= Rectangle{10.0, 10.0} // an instance of type Rectangle.

	t.Run("perimeter of rectangle", func(t *testing.T){
		got := rectangle.Perimeter()
		want := 40.0

		if want!=got{
			t.Errorf("got %.2f wanted %.2f", got, want) // here the %.2f is a placeholder where .2 means 2 decimal points
		}
	})
	t.Run("Perimeter of circle", func(t *testing.T){
		circle := Circle{10} // an instance of type Circle

		got := circle.Perimeter()
		want := 62.83185307179586

		if want != got{
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestArea(t *testing.T){
	areaTests:= [] struct{ // we have an array of structs, where every element is a shape and the expected value.
		shape Shape
		want float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, areaShape := range areaTests{
		got:= areaShape.shape.Area() // the area for that specifc shape.
		if got!= areaShape.want{
			t.Errorf("got %g want %g", got, areaShape.want)
		}
	}
}













