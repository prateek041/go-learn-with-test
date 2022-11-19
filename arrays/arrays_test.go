package arrays

import "testing"

func TestSumArray(t *testing.T){
	t.Run("Array of any size", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want:= 15
		assertMessageConstructor(t, got, want, numbers)
	})
}

func assertMessageConstructor(t testing.TB, got, want int, numbers []int){
	t.Helper()
	if got!=want{
		t.Errorf("got %d expected %d when passed %v", got, want, numbers)
	}
}
