//go:build !race

#include "textflag.h"

// func AtomicStoreU64Relaxed(dst *AtomicU64, src uint64)
TEXT ·AtomicStoreU64Relaxed(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64Relaxed(SB)

// func AtomicStoreU64Release(dst *AtomicU64, src uint64)
TEXT ·AtomicStoreU64Release(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64Release(SB)

// func AtomicStoreU64SeqCst(dst *AtomicU64, src uint64)
TEXT ·AtomicStoreU64SeqCst(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64SeqCst(SB)

// func AtomicStoreU64(dst *AtomicU64, src uint64)
TEXT ·AtomicStoreU64(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64SeqCst(SB)

// func AtomicLoadU64Relaxed(src *AtomicU64) uint64
TEXT ·AtomicLoadU64Relaxed(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64Relaxed(SB)

// func AtomicLoadU64Acquire(src *AtomicU64) uint64
TEXT ·AtomicLoadU64Acquire(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64Acquire(SB)

// func AtomicLoadU64SeqCst(src *AtomicU64) uint64
TEXT ·AtomicLoadU64SeqCst(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64SeqCst(SB)

// AtomicLoadU64(src *AtomicU64) uint64
TEXT ·AtomicLoadU64(SB), NOSPLIT, $0-16
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64SeqCst(SB)

// func AtomicAddU64Relaxed(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAddU64Relaxed(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64Relaxed(SB)

// func AtomicAddU64Acquire(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAddU64Acquire(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64Acquire(SB)

// func AtomicAddU64Release(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAddU64Release(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64Release(SB)

// func AtomicAddU64AcqRel(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAddU64AcqRel(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64AcqRel(SB)

// func AtomicAddU64SeqCst(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAddU64SeqCst(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64SeqCst(SB)

// func AtomicAddU64(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAddU64(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64SeqCst(SB)

// func AtomicOrU64Relaxed(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicOrU64Relaxed(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64Relaxed(SB)

// func AtomicOrU64Acquire(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicOrU64Acquire(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64Acquire(SB)

// func AtomicOrU64Release(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicOrU64Release(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64Release(SB)

// func AtomicOrU64AcqRel(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicOrU64AcqRel(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64AcqRel(SB)

// func AtomicOrU64SeqCst(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicOrU64SeqCst(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64SeqCst(SB)

// func AtomicOrU64(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicOrU64(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64SeqCst(SB)

// func AtomicAndU64Relaxed(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAndU64Relaxed(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64Relaxed(SB)

// func AtomicAndU64Acquire(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAndU64Acquire(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64Acquire(SB)

// func AtomicAndU64Release(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAndU64Release(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64Release(SB)

// func AtomicAndU64AcqRel(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAndU64AcqRel(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64AcqRel(SB)

// func AtomicAndU64SeqCst(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAndU64SeqCst(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64SeqCst(SB)

// func AtomicAndU64(dst *AtomicU64, val uint64) uint64
TEXT ·AtomicAndU64(SB), NOSPLIT, $0-24
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64SeqCst(SB)

// func AtomicCASU64(dst *AtomicU64, old uint64, new uint64) bool
TEXT ·AtomicCASU64(SB), NOSPLIT, $0-25
    JMP withoutforget∕go_atomic∕src∕common·AtomicCAS64(SB)
    