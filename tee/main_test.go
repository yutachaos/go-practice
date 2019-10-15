package tee

import (
	"testing"
)

func TestTee(t *testing.T) {
	input := make(chan string)
	output1, output2 := Tee(input)
	input <- "test"
	if <-output1 != <-output2 {
		t.Errorf("Assert error failed output1: %s output2: %s", <-output1, <-output2)
	}
}
