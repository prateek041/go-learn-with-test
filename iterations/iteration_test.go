package iterations

import "testing"

func TestIteration(t *testing.T){
	got := Repeat("a")
	want := "aaaaa"

	if got != want{
		t.Errorf("Expected %q but got %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}