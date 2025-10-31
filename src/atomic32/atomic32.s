//go:build !race

#include "textflag.h"

// func AtomicStore32Relaxed(dst *Atomic32, src int32)
TEXT ·AtomicStore32Relaxed(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore32Relaxed(SB)

// func AtomicStore32Release(dst *Atomic32, src int32)
TEXT ·AtomicStore32Release(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore32Release(SB)

// func AtomicStore32SeqCst(dst *Atomic32, src int32)
TEXT ·AtomicStore32SeqCst(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore32SeqCst(SB)

// func AtomicStore32(dst *Atomic32, src int32)
TEXT ·AtomicStore32(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore32SeqCst(SB)

// func AtomicLoad32Relaxed(src *Atomic32) int32
TEXT ·AtomicLoad32Relaxed(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad32Relaxed(SB)

// func AtomicLoad32Acquire(src *Atomic32) int32
TEXT ·AtomicLoad32Acquire(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad32Acquire(SB)

// func AtomicLoad32SeqCst(src *Atomic32) int32
TEXT ·AtomicLoad32SeqCst(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad32SeqCst(SB)

// AtomicLoad32(src *Atomic32) int32
TEXT ·AtomicLoad32(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad32SeqCst(SB)

// func AtomicAdd32Relaxed(dst *Atomic32, val int32) int32
TEXT ·AtomicAdd32Relaxed(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32Relaxed(SB)

// func AtomicAdd32Acquire(dst *Atomic32, val int32) int32
TEXT ·AtomicAdd32Acquire(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32Acquire(SB)

// func AtomicAdd32Release(dst *Atomic32, val int32) int32
TEXT ·AtomicAdd32Release(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32Release(SB)

// func AtomicAdd32AcqRel(dst *Atomic32, val int32) int32
TEXT ·AtomicAdd32AcqRel(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32AcqRel(SB)

// func AtomicAdd32SeqCst(dst *Atomic32, val int32) int32
TEXT ·AtomicAdd32SeqCst(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32SeqCst(SB)

// func AtomicAdd32(dst *Atomic32, val int32) int32
TEXT ·AtomicAdd32(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32SeqCst(SB)

// func AtomicOr32Relaxed(dst *Atomic32, val int32) int32
TEXT ·AtomicOr32Relaxed(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32Relaxed(SB)

// func AtomicOr32Acquire(dst *Atomic32, val int32) int32
TEXT ·AtomicOr32Acquire(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32Acquire(SB)

// func AtomicOr32Release(dst *Atomic32, val int32) int32
TEXT ·AtomicOr32Release(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32Release(SB)

// func AtomicOr32AcqRel(dst *Atomic32, val int32) int32
TEXT ·AtomicOr32AcqRel(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32AcqRel(SB)

// func AtomicOr32SeqCst(dst *Atomic32, val int32) int32
TEXT ·AtomicOr32SeqCst(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32SeqCst(SB)

// func AtomicOr32(dst *Atomic32, val int32) int32
TEXT ·AtomicOr32(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32SeqCst(SB)

// func AtomicAnd32Relaxed(dst *Atomic32, val int32) int32
TEXT ·AtomicAnd32Relaxed(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32Relaxed(SB)

// func AtomicAnd32Acquire(dst *Atomic32, val int32) int32
TEXT ·AtomicAnd32Acquire(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32Acquire(SB)

// func AtomicAnd32Release(dst *Atomic32, val int32) int32
TEXT ·AtomicAnd32Release(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32Release(SB)

// func AtomicAnd32AcqRel(dst *Atomic32, val int32) int32
TEXT ·AtomicAnd32AcqRel(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32AcqRel(SB)

// func AtomicAnd32SeqCst(dst *Atomic32, val int32) int32
TEXT ·AtomicAnd32SeqCst(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32SeqCst(SB)

// func AtomicAnd32(dst *Atomic32, val int32) int32
TEXT ·AtomicAnd32(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32SeqCst(SB)

// func AtomicCAS32(dst *Atomic32, old int32, new int32) bool
TEXT ·AtomicCAS32(SB), NOSPLIT, $0-17
    JMP withoutforget∕go_atomic∕src∕common·AtomicCAS32(SB)
    