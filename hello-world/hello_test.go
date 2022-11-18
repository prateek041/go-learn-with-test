package main

import "testing"
// the testing package is used for automated testing.
// any function starting with Test and argument (t *testing.T) is executed/

func TestHello(t *testing.T){ // t of type *testing.T is the hook, to use t.<method>
	t.Run("Saying hello to people with name", func(t *testing.T){
		got := Hello("nihal")
		want := "Hello, nihal"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("when the string is emply. return hello world", func(t *testing.T){
		got:= Hello("")
		want:= "Hello, world"

		if got!= want{
			t.Errorf("got %q want %q", got, want)
		}
	})
}