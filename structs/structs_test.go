package structs

import "testing"

type Rectangle struct {
	width float64
	height float64
}

func TestPerimeter(t *testing.T){
	var width float64 = 10.0 // this is general/test width and height of triangle
	var height float64 = 10.0
	t.Run("perimeter of rectangle", func(t *testing.T){
		got := Perimeter(width, height)
		want := 40.0

		if want!=got{
			t.Errorf("got %.2f wanted %.2f", got, want) // here the %.2f is a placeholder where .2 means 2 decimal points
		}
	})
	t.Run("Area of a rectangle", func(t *testing.T){
		got := Area(width, height)
		want := 100.0

		if want!= got{
			t.Errorf("got %.2f wanted %.2f", got, want)
		}
	})
}