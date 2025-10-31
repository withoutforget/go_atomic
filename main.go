package main

import (
	"fmt"
	"withoutforget/go_atomic/src/atomic64"
)

func main() {
	a := atomic64.AtomicU64{}
	a.Store(uint64(42))
	fmt.Println(a.Load())
}
