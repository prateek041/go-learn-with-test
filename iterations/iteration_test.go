package iterations

import "testing"

func TestIteration(t *testing.T){
	got := Repeat('a')
	want := "aaaaa"

	if got != want{
		t.Errof("Expected %q but got %q", got, want)
	}
}