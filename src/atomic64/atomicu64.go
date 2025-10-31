package atomic64

import _ "withoutforget/go_atomic/src/common"

type AtomicU64 struct {
	Data uint64
}

func AtomicStoreU64(dst *AtomicU64, src uint64)
func AtomicLoadU64(src *AtomicU64) AtomicU64
func AtomicAddU64(dst *AtomicU64, val uint64)
func AtomicOrU64(dst *AtomicU64, val uint64)
func AtomicAndU64(dst *AtomicU64, val uint64)
func AtomicCASU64(dst *AtomicU64, old uint64, new uint64) bool
func (a *AtomicU64) Store(src uint64) {
	AtomicStoreU64(a, src)
}

func (a *AtomicU64) Load() AtomicU64 {
	return AtomicLoadU64(a)
}

func (a *AtomicU64) Add(val uint64) {
	AtomicAddU64(a, val)
}

func (a *AtomicU64) Or(val uint64) {
	AtomicOrU64(a, val)
}

func (a *AtomicU64) And(val uint64) {
	AtomicAndU64(a, val)
}

func (a *AtomicU64) CAS(old, new uint64) bool {
	return AtomicCASU64(a, old, new)
}
