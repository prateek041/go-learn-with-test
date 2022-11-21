package structs

import "testing"

type Rectangle struct{ // this can be accesed through out the program.
	width float64
	height float64
}

func TestPerimeter(t *testing.T){
 	
	rectangle:= Rectangle{10.0, 10.0} // a variable rectangle, of type Rectangle

	t.Run("perimeter of rectangle", func(t *testing.T){
		got := Perimeter(rectangle)
		want := 40.0

		if want!=got{
			t.Errorf("got %.2f wanted %.2f", got, want) // here the %.2f is a placeholder where .2 means 2 decimal points
		}
	})
}

func TestArea(t *testing.T){

	rectangle:= Rectangle{10.0, 10.0} // a variable rectangle, of type Rectangle
	
	t.Run("Area of a rectangle", func(t *testing.T){
		got := Area(rectangle)
		want := 100.0

		if want!= got{
			t.Errorf("got %.2f wanted %.2f", got, want)
		}
	})
}