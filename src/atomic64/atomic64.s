//go:build !race

#include "textflag.h"

// func AtomicStore64(dst *Atomic64, src Atomic64)
TEXT ·AtomicStore64(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64(SB)

// AtomicLoad64(src *Atomic64) Atomic64
TEXT ·AtomicLoad64(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64(SB)

// func AtomicAdd64(dst *Atomic64, val int64)
TEXT ·AtomicAdd64(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64(SB)

// func AtomicOr64(dst *Atomic64, val int64)
TEXT ·AtomicOr64(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64(SB)

// func AtomicAnd64(dst *Atomic64, val int64)
TEXT ·AtomicAnd64(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64(SB)

// func AtomicCAS64(dst *Atomic64, old int64, new int64) bool
TEXT ·AtomicCAS64(SB), NOSPLIT, $0-17
    JMP withoutforget∕go_atomic∕src∕common·AtomicCAS64(SB)
    