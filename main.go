package main

type Atomic32 struct {
	Data int32
}

func AtomicStore32(dst *Atomic32, src Atomic32)
func AtomicLoad32(src *Atomic32) Atomic32
func AtomicAdd32(dst *Atomic32, val int32)

var a Atomic32

const T = 64
const I = 1_000_000

func main() {
	for range T {
		go func() {
			for range I {
				AtomicAdd32(&a, 1)
			}
		}()
	}

	for AtomicLoad32(&a).Data != T*I {
	}
}
