package atomic32

type Atomic32 struct {
	Data int32
}

func AtomicStore32(dst *Atomic32, src Atomic32)
func AtomicLoad32(src *Atomic32) Atomic32
func AtomicAdd32(dst *Atomic32, val int32)
func AtomicOr32(dst *Atomic32, val int32)
func AtomicAnd32(dst *Atomic32, val int32)
func AtomicCAS32(dst *Atomic32, old int32, new int32) bool
