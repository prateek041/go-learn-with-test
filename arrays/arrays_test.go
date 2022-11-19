package arrays

import (
	"testing"
	"reflect"
)

func TestSumArray(t *testing.T){
	t.Run("Array of any size", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want:= 15
		assertMessageConstructor(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T){
	t.Run("Sum of two different arrays", func(t *testing.T){
		firstArray := []int{1, 2}
		secondArray  := []int{0, 9}

		got := SumAll(firstArray, secondArray)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want){
			t.Errorf("wanted %v got %v", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T){
	t.Run("Sum of all tails", func(t *testing.T){
		got := SumAllTails([] int{1, 2}, [] int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want){
			t.Errorf("wanted %v got %v", want, got)
		}
	})
	t.Run("Sum of empty slice", func(t *testing.T){
		got:= SumAllTails([]int {}, []int {3, 4, 6})
		want:= []int{0, 9}

		if !reflect.DeepEqual(got, want){
			t.Errorf("Got %v wanted %v", got , want)
		}
	})
}

func assertMessageConstructor(t testing.TB, got, want int, numbers []int){
	t.Helper()
	if got!=want{
		t.Errorf("got %d expected %d when passed %v", got, want, numbers)
	}
}
