//go:build !race

#include "textflag.h"

// func AtomicStoreU64(dst *AtomicU64, src AtomicU64)
TEXT ·AtomicStoreU64(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore64(SB)

// AtomicLoadU64(src *AtomicU64) AtomicU64
TEXT ·AtomicLoadU64(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad64(SB)

// func AtomicAddU64(dst *AtomicU64, val uint64)
TEXT ·AtomicAddU64(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd64(SB)

// func AtomicOrU64(dst *AtomicU64, val uint64)
TEXT ·AtomicOrU64(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr64(SB)

// func AtomicAndU64(dst *AtomicU64, val uint64)
TEXT ·AtomicAndU64(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd64(SB)

// func AtomicCASU64(dst *AtomicU64, old uint64, new uint64) bool
TEXT ·AtomicCASU64(SB), NOSPLIT, $0-17
    JMP withoutforget∕go_atomic∕src∕common·AtomicCAS64(SB)
    