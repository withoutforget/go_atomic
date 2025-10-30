package atomic32

import _ "withoutforget/go_atomic/src/common"

type AtomicU32 struct {
	Data uint32
}

func AtomicStoreU32(dst *AtomicU32, src AtomicU32)
func AtomicLoadU32(src *AtomicU32) AtomicU32
func AtomicAddU32(dst *AtomicU32, val uint32)
func AtomicOrU32(dst *AtomicU32, val uint32)
func AtomicAndU32(dst *AtomicU32, val uint32)
func AtomicCASU32(dst *AtomicU32, old uint32, new uint32) bool
func (a *AtomicU32) Store(src AtomicU32) {
	AtomicStoreU32(a, src)
}

func (a *AtomicU32) Load() AtomicU32 {
	return AtomicLoadU32(a)
}

func (a *AtomicU32) Add(val uint32) {
	AtomicAddU32(a, val)
}

func (a *AtomicU32) Or(val uint32) {
	AtomicOrU32(a, val)
}

func (a *AtomicU32) And(val uint32) {
	AtomicAndU32(a, val)
}

func (a *AtomicU32) CAS(old, new uint32) bool {
	return AtomicCASU32(a, old, new)
}
