package atomic32

import _ "withoutforget/go_atomic/src/common"

type Atomic32 struct {
	Data int32
}

func AtomicStore32Relaxed(dst *Atomic32, src int32)
func AtomicStore32Release(dst *Atomic32, src int32)
func AtomicStore32SeqCst(dst *Atomic32, src int32)
func AtomicStore32(dst *Atomic32, src int32)

func AtomicLoad32Relaxed(src *Atomic32) int32
func AtomicLoad32Acquire(src *Atomic32) int32
func AtomicLoad32SeqCst(src *Atomic32) int32
func AtomicLoad32(src *Atomic32) int32

func AtomicAdd32Relaxed(dst *Atomic32, val int32) int32
func AtomicAdd32Acquire(dst *Atomic32, val int32) int32
func AtomicAdd32Release(dst *Atomic32, val int32) int32
func AtomicAdd32AcqRel(dst *Atomic32, val int32) int32
func AtomicAdd32SeqCst(dst *Atomic32, val int32) int32
func AtomicAdd32(dst *Atomic32, val int32) int32

func AtomicOr32Relaxed(dst *Atomic32, val int32) int32
func AtomicOr32Acquire(dst *Atomic32, val int32) int32
func AtomicOr32Release(dst *Atomic32, val int32) int32
func AtomicOr32AcqRel(dst *Atomic32, val int32) int32
func AtomicOr32SeqCst(dst *Atomic32, val int32) int32
func AtomicOr32(dst *Atomic32, val int32) int32

func AtomicAnd32Relaxed(dst *Atomic32, val int32) int32
func AtomicAnd32Acquire(dst *Atomic32, val int32) int32
func AtomicAnd32Release(dst *Atomic32, val int32) int32
func AtomicAnd32AcqRel(dst *Atomic32, val int32) int32
func AtomicAnd32SeqCst(dst *Atomic32, val int32) int32
func AtomicAnd32(dst *Atomic32, val int32) int32

func AtomicCAS32(dst *Atomic32, old int32, new int32) bool

func (a *Atomic32) Store(src int32) {
	AtomicStore32(a, src)
}

func (a *Atomic32) StoreRelaxed(src int32) {
	AtomicStore32Relaxed(a, src)
}

func (a *Atomic32) StoreRelease(src int32) {
	AtomicStore32Release(a, src)
}

func (a *Atomic32) StoreSeqCst(src int32) {
	AtomicStore32SeqCst(a, src)
}

func (a *Atomic32) Load() int32 {
	return AtomicLoad32(a)
}

func (a *Atomic32) LoadRelaxed() int32 {
	return AtomicLoad32Relaxed(a)
}

func (a *Atomic32) LoadAcquire() int32 {
	return AtomicLoad32Acquire(a)
}

func (a *Atomic32) LoadSeqCst() int32 {
	return AtomicLoad32SeqCst(a)
}

func (a *Atomic32) Add(val int32) int32 {
	return AtomicAdd32(a, val)
}

func (a *Atomic32) AddRelaxed(val int32) int32 {
	return AtomicAdd32Relaxed(a, val)
}

func (a *Atomic32) AddAcquire(val int32) int32 {
	return AtomicAdd32Acquire(a, val)
}

func (a *Atomic32) AddRelease(val int32) int32 {
	return AtomicAdd32Release(a, val)
}

func (a *Atomic32) AddAcqRel(val int32) int32 {
	return AtomicAdd32AcqRel(a, val)
}

func (a *Atomic32) AddSeqCst(val int32) int32 {
	return AtomicAdd32SeqCst(a, val)
}

func (a *Atomic32) Or(val int32) int32 {
	return AtomicOr32(a, val)
}

func (a *Atomic32) OrRelaxed(val int32) int32 {
	return AtomicOr32Relaxed(a, val)
}

func (a *Atomic32) OrAcquire(val int32) int32 {
	return AtomicOr32Acquire(a, val)
}

func (a *Atomic32) OrRelease(val int32) int32 {
	return AtomicOr32Release(a, val)
}

func (a *Atomic32) OrAcqRel(val int32) int32 {
	return AtomicOr32AcqRel(a, val)
}

func (a *Atomic32) OrSeqCst(val int32) int32 {
	return AtomicOr32SeqCst(a, val)
}

func (a *Atomic32) And(val int32) int32 {
	return AtomicAnd32(a, val)
}

func (a *Atomic32) AndRelaxed(val int32) int32 {
	return AtomicAnd32Relaxed(a, val)
}

func (a *Atomic32) AndAcquire(val int32) int32 {
	return AtomicAnd32Acquire(a, val)
}

func (a *Atomic32) AndRelease(val int32) int32 {
	return AtomicAnd32Release(a, val)
}

func (a *Atomic32) AndAcqRel(val int32) int32 {
	return AtomicAnd32AcqRel(a, val)
}

func (a *Atomic32) AndSeqCst(val int32) int32 {
	return AtomicAnd32SeqCst(a, val)
}

func (a *Atomic32) CAS(old, new int32) bool {
	return AtomicCAS32(a, old, new)
}
