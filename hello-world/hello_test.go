package main

import "testing"
// the testing package is used for automated testing.
// any function starting with Test and argument (t *testing.T) is executed/

func TestHello(t *testing.T){ // t of type *testing.T is the hook, to use t.<method>
	t.Run("Saying hello to people with name", func(t *testing.T){
		got := Hello("nihal")
		want := "Hello, nihal"

		assertCorrectMessage(t, got, want)
	})
	t.Run("when the string is emply. return hello world", func(t *testing.T){
		got:= Hello("")
		want:= "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper() // we add this to tell the test suit that it is a helper functiona. if it fails, the function call line is reported.
	if got!=want{
		t.Errorf("got %q wanted %q", got, want)
	}
}