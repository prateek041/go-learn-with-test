package integers

import "testing"

func TestAdder(t *testing.T){
	t.Run("Return the sum of two numbers", func(t *testing.T){
		got := Adder(2, 2)
		expected := 4

		if got!= expected{
			t.Errorf("Expected '%d' got '%d'", expected, got)
		}
	})
}