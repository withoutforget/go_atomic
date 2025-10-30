package main

import (
	"fmt"
	"withoutforget/go_atomic/src/atomic32"
)

func main() {
	a := atomic32.AtomicU32{}
	for range 10 {
		a.Add(1)
	}
	fmt.Println(a.Load())
}
