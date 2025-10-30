//go:build !race


#include "textflag.h"


// func AtomicStore32(dst *Atomic32, src Atomic32)
TEXT ·AtomicStore32(SB), NOSPLIT, $0
    JMP ∕src∕common·AtomicStore32(SB)

// func AtomicLoad32(src *Atomic32) Atomic32
TEXT ·AtomicLoad32(SB), NOSPLIT, $0
    JMP common·AtomicLoad32(SB)

// func AtomicAdd32(dst *Atomic32, val int32)
TEXT ·AtomicAdd32(SB), NOSPLIT, $0
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32(SB)

// func AtomicOr32(dst *Atomic32, val int32)
TEXT ·AtomicOr32(SB), NOSPLIT, $0
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32(SB)

// func AtomicAnd32(dst *Atomic32, val int32)
TEXT ·AtomicAnd32(SB), NOSPLIT, $0
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32(SB)

// func AtomicCAS32(dst *Atomic32, old *int32, new int32)
TEXT ·AtomicCAS32(SB), NOSPLIT, $0
    JMP withoutforget∕go_atomic∕src∕common·AtomicCAS32(SB)
