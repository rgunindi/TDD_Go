package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage:= func (t testing.TB,got,want string) {
		//Helper marks the calling function as a test helper function. When printing file and line information, that function will be skipped. Helper may be called simultaneously from multiple goroutines.
		t.Helper()
	if got!=want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("say hello with name",func(t *testing.T) {
		got:=Hello("John","")
		want:="hello John"
		assertCorrectMessage(t,got,want)	
	})
	t.Run("say hello, support empty value",func(t *testing.T) {
		got:=Hello("","")
		want:="hello world"
		assertCorrectMessage(t,got,want)	
	})
	t.Run("in Spanish",func(t *testing.T) {
		got:=Hello("Sara","Spanish")
		want:="Hola, Sara"
		assertCorrectMessage(t,got,want)
	})
	t.Run("in French",func(t *testing.T) {
		got:=Hello("Sarah","French")
		want:="Bonjour, Sarah"
		assertCorrectMessage(t,got,want)
	})
}