//go:build !race

#include "textflag.h"

// func AtomicStoreU32(dst *AtomicU32, src AtomicU32)
TEXT ·AtomicStoreU32(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicStore32(SB)

// AtomicLoadU32(src *AtomicU32) AtomicU32
TEXT ·AtomicLoadU32(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicLoad32(SB)

// func AtomicAddU32(dst *AtomicU32, val uint32)
TEXT ·AtomicAddU32(SB), NOSPLIT, $0-20
    JMP withoutforget∕go_atomic∕src∕common·AtomicAdd32(SB)

// func AtomicOrU32(dst *AtomicU32, val uint32)
TEXT ·AtomicOrU32(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicOr32(SB)

// func AtomicAndU32(dst *AtomicU32, val uint32)
TEXT ·AtomicAndU32(SB), NOSPLIT, $0-12
    JMP withoutforget∕go_atomic∕src∕common·AtomicAnd32(SB)

// func AtomicCASU32(dst *AtomicU32, old uint32, new uint32) bool
TEXT ·AtomicCASU32(SB), NOSPLIT, $0-17
    JMP withoutforget∕go_atomic∕src∕common·AtomicCAS32(SB)
    