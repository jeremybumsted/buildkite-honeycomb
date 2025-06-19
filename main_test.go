package main

import "testing"

func TestAdd(t *testing.T){
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSubtract(t *testing.T){
	got := Subtract(5,2)
	want := 3

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMultiply(t *testing.T){
	got := Multiply(5,2)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDivide(t *testing.T){
	got := Divide(5,2)
	want := 2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRandomFailure(t *testing.T) {
	got := RandomBoolean()
	want := true
	
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}