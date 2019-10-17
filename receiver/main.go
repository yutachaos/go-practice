package main

import (
	"fmt"
)

type testFunc func(string)
type test testFunc

func main() {
	var t test = func(arg string) {
		fmt.Printf("arg: %s\n", arg)
	}
	t("test")
	t.call("test")
}

func (t test) call(arg1 string) {
	t(arg1)
}
