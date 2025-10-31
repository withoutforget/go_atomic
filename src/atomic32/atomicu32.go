package atomic32

import _ "withoutforget/go_atomic/src/common"

type AtomicU32 struct {
	Data uint32
}

func AtomicStoreU32Relaxed(dst *AtomicU32, src uint32)
func AtomicStoreU32Release(dst *AtomicU32, src uint32)
func AtomicStoreU32SeqCst(dst *AtomicU32, src uint32)
func AtomicStoreU32(dst *AtomicU32, src uint32)

func AtomicLoadU32Relaxed(src *AtomicU32) uint32
func AtomicLoadU32Acquire(src *AtomicU32) uint32
func AtomicLoadU32SeqCst(src *AtomicU32) uint32
func AtomicLoadU32(src *AtomicU32) uint32

func AtomicAddU32Relaxed(dst *AtomicU32, val uint32) uint32
func AtomicAddU32Acquire(dst *AtomicU32, val uint32) uint32
func AtomicAddU32Release(dst *AtomicU32, val uint32) uint32
func AtomicAddU32AcqRel(dst *AtomicU32, val uint32) uint32
func AtomicAddU32SeqCst(dst *AtomicU32, val uint32) uint32
func AtomicAddU32(dst *AtomicU32, val uint32) uint32

func AtomicOrU32Relaxed(dst *AtomicU32, val uint32) uint32
func AtomicOrU32Acquire(dst *AtomicU32, val uint32) uint32
func AtomicOrU32Release(dst *AtomicU32, val uint32) uint32
func AtomicOrU32AcqRel(dst *AtomicU32, val uint32) uint32
func AtomicOrU32SeqCst(dst *AtomicU32, val uint32) uint32
func AtomicOrU32(dst *AtomicU32, val uint32) uint32

func AtomicAndU32Relaxed(dst *AtomicU32, val uint32) uint32
func AtomicAndU32Acquire(dst *AtomicU32, val uint32) uint32
func AtomicAndU32Release(dst *AtomicU32, val uint32) uint32
func AtomicAndU32AcqRel(dst *AtomicU32, val uint32) uint32
func AtomicAndU32SeqCst(dst *AtomicU32, val uint32) uint32
func AtomicAndU32(dst *AtomicU32, val uint32) uint32

func AtomicCASU32(dst *AtomicU32, old uint32, new uint32) bool

func (a *AtomicU32) Store(src uint32) {
	AtomicStoreU32(a, src)
}

func (a *AtomicU32) StoreRelaxed(src uint32) {
	AtomicStoreU32Relaxed(a, src)
}

func (a *AtomicU32) StoreRelease(src uint32) {
	AtomicStoreU32Release(a, src)
}

func (a *AtomicU32) StoreSeqCst(src uint32) {
	AtomicStoreU32SeqCst(a, src)
}

func (a *AtomicU32) Load() uint32 {
	return AtomicLoadU32(a)
}

func (a *AtomicU32) LoadRelaxed() uint32 {
	return AtomicLoadU32Relaxed(a)
}

func (a *AtomicU32) LoadAcquire() uint32 {
	return AtomicLoadU32Acquire(a)
}

func (a *AtomicU32) LoadSeqCst() uint32 {
	return AtomicLoadU32SeqCst(a)
}

func (a *AtomicU32) Add(val uint32) uint32 {
	return AtomicAddU32(a, val)
}

func (a *AtomicU32) AddRelaxed(val uint32) uint32 {
	return AtomicAddU32Relaxed(a, val)
}

func (a *AtomicU32) AddAcquire(val uint32) uint32 {
	return AtomicAddU32Acquire(a, val)
}

func (a *AtomicU32) AddRelease(val uint32) uint32 {
	return AtomicAddU32Release(a, val)
}

func (a *AtomicU32) AddAcqRel(val uint32) uint32 {
	return AtomicAddU32AcqRel(a, val)
}

func (a *AtomicU32) AddSeqCst(val uint32) uint32 {
	return AtomicAddU32SeqCst(a, val)
}

func (a *AtomicU32) Or(val uint32) uint32 {
	return AtomicOrU32(a, val)
}

func (a *AtomicU32) OrRelaxed(val uint32) uint32 {
	return AtomicOrU32Relaxed(a, val)
}

func (a *AtomicU32) OrAcquire(val uint32) uint32 {
	return AtomicOrU32Acquire(a, val)
}

func (a *AtomicU32) OrRelease(val uint32) uint32 {
	return AtomicOrU32Release(a, val)
}

func (a *AtomicU32) OrAcqRel(val uint32) uint32 {
	return AtomicOrU32AcqRel(a, val)
}

func (a *AtomicU32) OrSeqCst(val uint32) uint32 {
	return AtomicOrU32SeqCst(a, val)
}

func (a *AtomicU32) And(val uint32) uint32 {
	return AtomicAndU32(a, val)
}

func (a *AtomicU32) AndRelaxed(val uint32) uint32 {
	return AtomicAndU32Relaxed(a, val)
}

func (a *AtomicU32) AndAcquire(val uint32) uint32 {
	return AtomicAndU32Acquire(a, val)
}

func (a *AtomicU32) AndRelease(val uint32) uint32 {
	return AtomicAndU32Release(a, val)
}

func (a *AtomicU32) AndAcqRel(val uint32) uint32 {
	return AtomicAndU32AcqRel(a, val)
}

func (a *AtomicU32) AndSeqCst(val uint32) uint32 {
	return AtomicAndU32SeqCst(a, val)
}

func (a *AtomicU32) CAS(old, new uint32) bool {
	return AtomicCASU32(a, old, new)
}
