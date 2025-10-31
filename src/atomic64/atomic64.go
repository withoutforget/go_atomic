package atomic64

import _ "withoutforget/go_atomic/src/common"

type Atomic64 struct {
	Data int64
}

func AtomicStore64(dst *Atomic64, src int64)
func AtomicLoad64(src *Atomic64) Atomic64
func AtomicAdd64(dst *Atomic64, val int64)
func AtomicOr64(dst *Atomic64, val int64)
func AtomicAnd64(dst *Atomic64, val int64)
func AtomicCAS64(dst *Atomic64, old int64, new int64) bool

func (a *Atomic64) Store(src int64) {
	AtomicStore64(a, src)
}
func (a *Atomic64) Load() Atomic64 {
	return AtomicLoad64(a)
}

func (a *Atomic64) Add(val int64) {
	AtomicAdd64(a, val)
}

func (a *Atomic64) Or(val int64) {
	AtomicOr64(a, val)
}

func (a *Atomic64) And(val int64) {
	AtomicAnd64(a, val)
}

func (a *Atomic64) CAS(old, new int64) bool {
	return AtomicCAS64(a, old, new)
}
