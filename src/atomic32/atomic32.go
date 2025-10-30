package atomic32

import _ "withoutforget/go_atomic/src/common"

type Atomic32 struct {
	Data int32
}

func AtomicStore32(dst *Atomic32, src Atomic32)
func AtomicLoad32(src *Atomic32) Atomic32
func AtomicAdd32(dst *Atomic32, val int32)
func AtomicOr32(dst *Atomic32, val int32)
func AtomicAnd32(dst *Atomic32, val int32)
func AtomicCAS32(dst *Atomic32, old int32, new int32) bool

func (a *Atomic32) Store(src Atomic32) {
	AtomicStore32(a, src)
}
func (a *Atomic32) Load() Atomic32 {
	return AtomicLoad32(a)
}

func (a *Atomic32) Add(val int32) {
	AtomicAdd32(a, val)
}

func (a *Atomic32) Or(val int32) {
	AtomicOr32(a, val)
}

func (a *Atomic32) And(val int32) {
	AtomicAnd32(a, val)
}

func (a *Atomic32) CAS(old, new int32) bool {
	return AtomicCAS32(a, old, new)
}
