package main

import (
	"fmt"
)

type foobar interface {
	foo()
	bar()
}

type itemA struct{}

func (a *itemA) foo() {
	fmt.Println("foo on A")
}

func (a *itemA) bar() {
	fmt.Println("bar on A")
}

type itemB struct{}

func (a *itemB) foo() {
	fmt.Println("foo on B")
}

func (a *itemB) bar() {
	fmt.Println("bar on B")
}

func doFoo(item foobar) {
	item.foo()
}

func doBar(item foobar) {
	item.bar()
}

func main() {
	doFoo(&itemA{}) // Prints "foo on A"
	doFoo(&itemB{}) // Prints "foo on B"
}
