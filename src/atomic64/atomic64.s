//go:build !race

#include "textflag.h"

// func AtomicStore64Relaxed(dst *Atomic64, src int64)
TEXT ·AtomicStore64Relaxed(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64Relaxed(SB)

// func AtomicStore64Release(dst *Atomic64, src int64)
TEXT ·AtomicStore64Release(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64Release(SB)

// func AtomicStore64SeqCst(dst *Atomic64, src int64)
TEXT ·AtomicStore64SeqCst(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64SeqCst(SB)

// func AtomicStore64(dst *Atomic64, src int64)
TEXT ·AtomicStore64(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64SeqCst(SB)

// func AtomicLoad64Relaxed(src *Atomic64) int64
TEXT ·AtomicLoad64Relaxed(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64Relaxed(SB)

// func AtomicLoad64Acquire(src *Atomic64) int64
TEXT ·AtomicLoad64Acquire(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64Acquire(SB)

// func AtomicLoad64SeqCst(src *Atomic64) int64
TEXT ·AtomicLoad64SeqCst(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64SeqCst(SB)

// AtomicLoad64(src *Atomic64) int64
TEXT ·AtomicLoad64(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64SeqCst(SB)

// func AtomicAdd64Relaxed(dst *Atomic64, val int64) int64
TEXT ·AtomicAdd64Relaxed(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64Relaxed(SB)

// func AtomicAdd64Acquire(dst *Atomic64, val int64) int64
TEXT ·AtomicAdd64Acquire(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64Acquire(SB)

// func AtomicAdd64Release(dst *Atomic64, val int64) int64
TEXT ·AtomicAdd64Release(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64Release(SB)

// func AtomicAdd64AcqRel(dst *Atomic64, val int64) int64
TEXT ·AtomicAdd64AcqRel(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64AcqRel(SB)

// func AtomicAdd64SeqCst(dst *Atomic64, val int64) int64
TEXT ·AtomicAdd64SeqCst(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64SeqCst(SB)

// func AtomicAdd64(dst *Atomic64, val int64) int64
TEXT ·AtomicAdd64(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64SeqCst(SB)

// func AtomicOr64Relaxed(dst *Atomic64, val int64) int64
TEXT ·AtomicOr64Relaxed(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64Relaxed(SB)

// func AtomicOr64Acquire(dst *Atomic64, val int64) int64
TEXT ·AtomicOr64Acquire(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64Acquire(SB)

// func AtomicOr64Release(dst *Atomic64, val int64) int64
TEXT ·AtomicOr64Release(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64Release(SB)

// func AtomicOr64AcqRel(dst *Atomic64, val int64) int64
TEXT ·AtomicOr64AcqRel(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64AcqRel(SB)

// func AtomicOr64SeqCst(dst *Atomic64, val int64) int64
TEXT ·AtomicOr64SeqCst(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64SeqCst(SB)

// func AtomicOr64(dst *Atomic64, val int64) int64
TEXT ·AtomicOr64(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64SeqCst(SB)

// func AtomicAnd64Relaxed(dst *Atomic64, val int64) int64
TEXT ·AtomicAnd64Relaxed(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64Relaxed(SB)

// func AtomicAnd64Acquire(dst *Atomic64, val int64) int64
TEXT ·AtomicAnd64Acquire(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64Acquire(SB)

// func AtomicAnd64Release(dst *Atomic64, val int64) int64
TEXT ·AtomicAnd64Release(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64Release(SB)

// func AtomicAnd64AcqRel(dst *Atomic64, val int64) int64
TEXT ·AtomicAnd64AcqRel(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64AcqRel(SB)

// func AtomicAnd64SeqCst(dst *Atomic64, val int64) int64
TEXT ·AtomicAnd64SeqCst(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64SeqCst(SB)

// func AtomicAnd64(dst *Atomic64, val int64) int64
TEXT ·AtomicAnd64(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64SeqCst(SB)

// func AtomicCAS64(dst *Atomic64, old int64, new int64) bool
TEXT ·AtomicCAS64(SB), NOSPLIT, $0-25
    JMP withoutforget∕go_atomic∕src∕common·AtomicCAS64(SB)
    