package main

import (
	"fmt"
	atomic32 "withoutforget/go_atomic/src/atomic32"
)

func main() {
	a := atomic32.Atomic32{}
	atomic32.AtomicAdd32(&a, 32)
	fmt.Println(atomic32.AtomicLoad32(&a))
}
