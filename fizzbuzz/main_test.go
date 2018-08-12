package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	actual1 := FizzBuzz(1)
	expected1 := "1"
	if expected1 != actual1 {
		t.Fatal("Test fail")
	}

	actual2 := FizzBuzz(3)
	expected2 := "Fizz"
	if expected2 != actual2 {
		t.Fatal("Test fail")
	}

	actual3 := FizzBuzz(5)
	expected3 := "Buzz"
	if expected3 != actual3 {
		t.Fatal("Test fail")
	}

	actual4 := FizzBuzz(15)
	expected4 := "FizzBuzz"
	if expected4 != actual4 {
		t.Fatal("Test fail")
	}
}
