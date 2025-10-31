//go:build !race

#include "textflag.h"

// func AtomicStoreU32Relaxed(dst *AtomicU32, src uint32)
TEXT ·AtomicStoreU32Relaxed(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore32Relaxed(SB)

// func AtomicStoreU32Release(dst *AtomicU32, src uint32)
TEXT ·AtomicStoreU32Release(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore32Release(SB)

// func AtomicStoreU32SeqCst(dst *AtomicU32, src uint32)
TEXT ·AtomicStoreU32SeqCst(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore32SeqCst(SB)

// func AtomicStoreU32(dst *AtomicU32, src uint32)
TEXT ·AtomicStoreU32(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore32SeqCst(SB)

// func AtomicLoadU32Relaxed(src *AtomicU32) uint32
TEXT ·AtomicLoadU32Relaxed(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad32Relaxed(SB)

// func AtomicLoadU32Acquire(src *AtomicU32) uint32
TEXT ·AtomicLoadU32Acquire(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad32Acquire(SB)

// func AtomicLoadU32SeqCst(src *AtomicU32) uint32
TEXT ·AtomicLoadU32SeqCst(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad32SeqCst(SB)

// AtomicLoadU32(src *AtomicU32) uint32
TEXT ·AtomicLoadU32(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad32SeqCst(SB)

// func AtomicAddU32Relaxed(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAddU32Relaxed(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32Relaxed(SB)

// func AtomicAddU32Acquire(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAddU32Acquire(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32Acquire(SB)

// func AtomicAddU32Release(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAddU32Release(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32Release(SB)

// func AtomicAddU32AcqRel(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAddU32AcqRel(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32AcqRel(SB)

// func AtomicAddU32SeqCst(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAddU32SeqCst(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32SeqCst(SB)

// func AtomicAddU32(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAddU32(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32SeqCst(SB)

// func AtomicOrU32Relaxed(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicOrU32Relaxed(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32Relaxed(SB)

// func AtomicOrU32Acquire(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicOrU32Acquire(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32Acquire(SB)

// func AtomicOrU32Release(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicOrU32Release(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32Release(SB)

// func AtomicOrU32AcqRel(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicOrU32AcqRel(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32AcqRel(SB)

// func AtomicOrU32SeqCst(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicOrU32SeqCst(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32SeqCst(SB)

// func AtomicOrU32(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicOrU32(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32SeqCst(SB)

// func AtomicAndU32Relaxed(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAndU32Relaxed(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32Relaxed(SB)

// func AtomicAndU32Acquire(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAndU32Acquire(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32Acquire(SB)

// func AtomicAndU32Release(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAndU32Release(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32Release(SB)

// func AtomicAndU32AcqRel(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAndU32AcqRel(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32AcqRel(SB)

// func AtomicAndU32SeqCst(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAndU32SeqCst(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32SeqCst(SB)

// func AtomicAndU32(dst *AtomicU32, val uint32) uint32
TEXT ·AtomicAndU32(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32SeqCst(SB)

// func AtomicCASU32(dst *AtomicU32, old uint32, new uint32) bool
TEXT ·AtomicCASU32(SB), NOSPLIT, $0-17
    JMP withoutforget∕go_atomic∕src∕common·AtomicCAS32(SB)
    