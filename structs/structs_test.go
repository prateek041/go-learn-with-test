package structs

import "testing"

func TestPerimeter(t *testing.T){
	t.Run("perimeter of rectangle", func(t *testing.T){
		var width float64 = 10.0
		var height float64 = 10.0
		got := Perimeter(width, height)
		want := 40.0

		if(want!=got){
			t.Errorf("got %.2f wanted %.2f", got, want) // here the %.2f is a placeholder where .2 means 2 decimal points
		}
	})
}