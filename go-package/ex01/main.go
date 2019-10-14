package main

import (
	"fmt"
	"github.com/yutachaos/go-practice/go-package/ex01/sample1"
)

func main() {
	// これはアクセスできる
	fmt.Printf("Foo: %s", sample1.Foo)
	fmt.Printf("Hoge: %s", sample1.Hoge)

	// これはアクセス出来ない
	fmt.Printf("bar: %s", sample1.bar)
	fmt.Printf("fuga: %s", sample1.fuga)
}
