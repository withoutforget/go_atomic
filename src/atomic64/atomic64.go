package atomic64

import _ "withoutforget/go_atomic/src/common"

type Atomic64 struct {
	Data int64
}

func AtomicStore64Relaxed(dst *Atomic64, src int64)
func AtomicStore64Release(dst *Atomic64, src int64)
func AtomicStore64SeqCst(dst *Atomic64, src int64)
func AtomicStore64(dst *Atomic64, src int64)

func AtomicLoad64Relaxed(src *Atomic64) int64
func AtomicLoad64Acquire(src *Atomic64) int64
func AtomicLoad64SeqCst(src *Atomic64) int64
func AtomicLoad64(src *Atomic64) int64

func AtomicAdd64Relaxed(dst *Atomic64, val int64) int64
func AtomicAdd64Acquire(dst *Atomic64, val int64) int64
func AtomicAdd64Release(dst *Atomic64, val int64) int64
func AtomicAdd64AcqRel(dst *Atomic64, val int64) int64
func AtomicAdd64SeqCst(dst *Atomic64, val int64) int64
func AtomicAdd64(dst *Atomic64, val int64) int64

func AtomicOr64Relaxed(dst *Atomic64, val int64) int64
func AtomicOr64Acquire(dst *Atomic64, val int64) int64
func AtomicOr64Release(dst *Atomic64, val int64) int64
func AtomicOr64AcqRel(dst *Atomic64, val int64) int64
func AtomicOr64SeqCst(dst *Atomic64, val int64) int64
func AtomicOr64(dst *Atomic64, val int64) int64

func AtomicAnd64Relaxed(dst *Atomic64, val int64) int64
func AtomicAnd64Acquire(dst *Atomic64, val int64) int64
func AtomicAnd64Release(dst *Atomic64, val int64) int64
func AtomicAnd64AcqRel(dst *Atomic64, val int64) int64
func AtomicAnd64SeqCst(dst *Atomic64, val int64) int64
func AtomicAnd64(dst *Atomic64, val int64) int64

func AtomicCAS64(dst *Atomic64, old int64, new int64) bool

func (a *Atomic64) Store(src int64) {
	AtomicStore64(a, src)
}

func (a *Atomic64) StoreRelaxed(src int64) {
	AtomicStore64Relaxed(a, src)
}

func (a *Atomic64) StoreRelease(src int64) {
	AtomicStore64Release(a, src)
}

func (a *Atomic64) StoreSeqCst(src int64) {
	AtomicStore64SeqCst(a, src)
}

func (a *Atomic64) Load() int64 {
	return AtomicLoad64(a)
}

func (a *Atomic64) LoadRelaxed() int64 {
	return AtomicLoad64Relaxed(a)
}

func (a *Atomic64) LoadAcquire() int64 {
	return AtomicLoad64Acquire(a)
}

func (a *Atomic64) LoadSeqCst() int64 {
	return AtomicLoad64SeqCst(a)
}

// Add methods
func (a *Atomic64) Add(val int64) int64 {
	return AtomicAdd64(a, val)
}

func (a *Atomic64) AddRelaxed(val int64) int64 {
	return AtomicAdd64Relaxed(a, val)
}

func (a *Atomic64) AddAcquire(val int64) int64 {
	return AtomicAdd64Acquire(a, val)
}

func (a *Atomic64) AddRelease(val int64) int64 {
	return AtomicAdd64Release(a, val)
}

func (a *Atomic64) AddAcqRel(val int64) int64 {
	return AtomicAdd64AcqRel(a, val)
}

func (a *Atomic64) AddSeqCst(val int64) int64 {
	return AtomicAdd64SeqCst(a, val)
}

func (a *Atomic64) Or(val int64) int64 {
	return AtomicOr64(a, val)
}

func (a *Atomic64) OrRelaxed(val int64) int64 {
	return AtomicOr64Relaxed(a, val)
}

func (a *Atomic64) OrAcquire(val int64) int64 {
	return AtomicOr64Acquire(a, val)
}

func (a *Atomic64) OrRelease(val int64) int64 {
	return AtomicOr64Release(a, val)
}

func (a *Atomic64) OrAcqRel(val int64) int64 {
	return AtomicOr64AcqRel(a, val)
}

func (a *Atomic64) OrSeqCst(val int64) int64 {
	return AtomicOr64SeqCst(a, val)
}

func (a *Atomic64) And(val int64) int64 {
	return AtomicAnd64(a, val)
}

func (a *Atomic64) AndRelaxed(val int64) int64 {
	return AtomicAnd64Relaxed(a, val)
}

func (a *Atomic64) AndAcquire(val int64) int64 {
	return AtomicAnd64Acquire(a, val)
}

func (a *Atomic64) AndRelease(val int64) int64 {
	return AtomicAnd64Release(a, val)
}

func (a *Atomic64) AndAcqRel(val int64) int64 {
	return AtomicAnd64AcqRel(a, val)
}

func (a *Atomic64) AndSeqCst(val int64) int64 {
	return AtomicAnd64SeqCst(a, val)
}

func (a *Atomic64) CAS(old, new int64) bool {
	return AtomicCAS64(a, old, new)
}
