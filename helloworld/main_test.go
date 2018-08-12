package main

import "testing"

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld()
	expected := "Hello Go World!!!!!"
	if expected != actual {
		t.Fatal("Test fail")
	}
}
