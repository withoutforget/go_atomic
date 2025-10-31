package atomic64

import _ "withoutforget/go_atomic/src/common"

type AtomicU64 struct {
	Data uint64
}

func AtomicStoreU64Relaxed(dst *AtomicU64, src uint64)
func AtomicStoreU64Release(dst *AtomicU64, src uint64)
func AtomicStoreU64SeqCst(dst *AtomicU64, src uint64)
func AtomicStoreU64(dst *AtomicU64, src uint64)

func AtomicLoadU64Relaxed(src *AtomicU64) uint64
func AtomicLoadU64Acquire(src *AtomicU64) uint64
func AtomicLoadU64SeqCst(src *AtomicU64) uint64
func AtomicLoadU64(src *AtomicU64) uint64

func AtomicAddU64Relaxed(dst *AtomicU64, val uint64) uint64
func AtomicAddU64Acquire(dst *AtomicU64, val uint64) uint64
func AtomicAddU64Release(dst *AtomicU64, val uint64) uint64
func AtomicAddU64AcqRel(dst *AtomicU64, val uint64) uint64
func AtomicAddU64SeqCst(dst *AtomicU64, val uint64) uint64
func AtomicAddU64(dst *AtomicU64, val uint64) uint64

func AtomicOrU64Relaxed(dst *AtomicU64, val uint64) uint64
func AtomicOrU64Acquire(dst *AtomicU64, val uint64) uint64
func AtomicOrU64Release(dst *AtomicU64, val uint64) uint64
func AtomicOrU64AcqRel(dst *AtomicU64, val uint64) uint64
func AtomicOrU64SeqCst(dst *AtomicU64, val uint64) uint64
func AtomicOrU64(dst *AtomicU64, val uint64) uint64

func AtomicAndU64Relaxed(dst *AtomicU64, val uint64) uint64
func AtomicAndU64Acquire(dst *AtomicU64, val uint64) uint64
func AtomicAndU64Release(dst *AtomicU64, val uint64) uint64
func AtomicAndU64AcqRel(dst *AtomicU64, val uint64) uint64
func AtomicAndU64SeqCst(dst *AtomicU64, val uint64) uint64
func AtomicAndU64(dst *AtomicU64, val uint64) uint64

func AtomicCASU64(dst *AtomicU64, old uint64, new uint64) bool

func (a *AtomicU64) Store(src uint64) {
	AtomicStoreU64(a, src)
}

func (a *AtomicU64) StoreRelaxed(src uint64) {
	AtomicStoreU64Relaxed(a, src)
}

func (a *AtomicU64) StoreRelease(src uint64) {
	AtomicStoreU64Release(a, src)
}

func (a *AtomicU64) StoreSeqCst(src uint64) {
	AtomicStoreU64SeqCst(a, src)
}

func (a *AtomicU64) Load() uint64 {
	return AtomicLoadU64(a)
}

func (a *AtomicU64) LoadRelaxed() uint64 {
	return AtomicLoadU64Relaxed(a)
}

func (a *AtomicU64) LoadAcquire() uint64 {
	return AtomicLoadU64Acquire(a)
}

func (a *AtomicU64) LoadSeqCst() uint64 {
	return AtomicLoadU64SeqCst(a)
}

func (a *AtomicU64) Add(val uint64) uint64 {
	return AtomicAddU64(a, val)
}

func (a *AtomicU64) AddRelaxed(val uint64) uint64 {
	return AtomicAddU64Relaxed(a, val)
}

func (a *AtomicU64) AddAcquire(val uint64) uint64 {
	return AtomicAddU64Acquire(a, val)
}

func (a *AtomicU64) AddRelease(val uint64) uint64 {
	return AtomicAddU64Release(a, val)
}

func (a *AtomicU64) AddAcqRel(val uint64) uint64 {
	return AtomicAddU64AcqRel(a, val)
}

func (a *AtomicU64) AddSeqCst(val uint64) uint64 {
	return AtomicAddU64SeqCst(a, val)
}

func (a *AtomicU64) Or(val uint64) uint64 {
	return AtomicOrU64(a, val)
}

func (a *AtomicU64) OrRelaxed(val uint64) uint64 {
	return AtomicOrU64Relaxed(a, val)
}

func (a *AtomicU64) OrAcquire(val uint64) uint64 {
	return AtomicOrU64Acquire(a, val)
}

func (a *AtomicU64) OrRelease(val uint64) uint64 {
	return AtomicOrU64Release(a, val)
}

func (a *AtomicU64) OrAcqRel(val uint64) uint64 {
	return AtomicOrU64AcqRel(a, val)
}

func (a *AtomicU64) OrSeqCst(val uint64) uint64 {
	return AtomicOrU64SeqCst(a, val)
}

func (a *AtomicU64) And(val uint64) uint64 {
	return AtomicAndU64(a, val)
}

func (a *AtomicU64) AndRelaxed(val uint64) uint64 {
	return AtomicAndU64Relaxed(a, val)
}

func (a *AtomicU64) AndAcquire(val uint64) uint64 {
	return AtomicAndU64Acquire(a, val)
}

func (a *AtomicU64) AndRelease(val uint64) uint64 {
	return AtomicAndU64Release(a, val)
}

func (a *AtomicU64) AndAcqRel(val uint64) uint64 {
	return AtomicAndU64AcqRel(a, val)
}

func (a *AtomicU64) AndSeqCst(val uint64) uint64 {
	return AtomicAndU64SeqCst(a, val)
}

func (a *AtomicU64) CAS(old, new uint64) bool {
	return AtomicCASU64(a, old, new)
}
