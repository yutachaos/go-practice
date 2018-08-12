package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	actual := FizzBuzz(1)
	expected := "1"
	if expected != actual {
		t.Fatal("Test fail")
	}
}
