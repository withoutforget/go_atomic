package main

import (
	"fmt"
	"withoutforget/go_atomic/src/atomic64"
)

func main() {
	a := atomic64.Atomic64{}
	b := a.OrRelaxed(0b1010)
	c := a.OrRelaxed(0b0101)
	fmt.Printf("%b\n%b\n%b\n", a.Load(), b, c)
}
