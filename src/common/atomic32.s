#include "textflag.h"

// func AtomicStore32Relaxed(dst *int32, src int32)
TEXT ·AtomicStore32Relaxed(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    XCHGL BX, (AX)
    RET

// func AtomicStore32Release(dst *int32, src int32)
TEXT ·AtomicStore32Release(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    MOVL BX, (AX)
    RET

// func AtomicStore32SeqCst(dst *int32, src int32)
TEXT ·AtomicStore32SeqCst(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    MOVL BX, (AX)
    RET


// func AtomicLoad32Relaxed(src *int32) int32
TEXT ·AtomicLoad32Relaxed(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL (AX), AX
    XCHGL AX, ret+8(FP)
    RET

// func AtomicLoad32Acquire(src *int32) int32
TEXT ·AtomicLoad32Acquire(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL (AX), AX
    MOVL AX, ret+8(FP)
    RET

// func AtomicLoad32SeqCst(src *int32) int32
TEXT ·AtomicLoad32SeqCst(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL (AX), AX
    MOVL AX, ret+8(FP)
    RET

// func AtomicAdd32Relaxed(dst *int32, val int32)
TEXT ·AtomicAdd32Relaxed(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    XADDL BX, (AX)
    RET
// func AtomicAdd32Acquire(dst *int32, val int32)
TEXT ·AtomicAdd32Acquire(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    XADDL BX, (AX)
    RET
// func AtomicAdd32Release(dst *int32, val int32)
TEXT ·AtomicAdd32Release(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    XADDL BX, (AX)
    RET
// func AtomicAdd32AcqRel(dst *int32, val int32)
TEXT ·AtomicAdd32AcqRel(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    XADDL BX, (AX)
    RET
// func AtomicAdd32SeqCst(dst *int32, val int32)
TEXT ·AtomicAdd32SeqCst(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    XADDL BX, (AX)
    RET


// func AtomicOr32Relaxed(dst *int32, val int32)
TEXT ·AtomicOr32Relaxed(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ORL BX, (AX)
    RET

// func AtomicOr32Acquire(dst *int32, val int32)
TEXT ·AtomicOr32Acquire(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ORL BX, (AX)
    RET

// func AtomicOr32Release(dst *int32, val int32)
TEXT ·AtomicOr32Release(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ORL BX, (AX)
    RET

// func AtomicOr32AcqRel(dst *int32, val int32)
TEXT ·AtomicOr32AcqRel(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ORL BX, (AX)
    RET

// func AtomicOr32SeqCst(dst *int32, val int32)
TEXT ·AtomicOr32SeqCst(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ORL BX, (AX)
    RET

// func AtomicAnd32Relaxed(dst *int32, val int32)
TEXT ·AtomicAnd32Relaxed(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ANDL BX, (AX)
    RET
// func AtomicAnd32Acquire(dst *int32, val int32)
TEXT ·AtomicAnd32Acquire(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ANDL BX, (AX)
    RET
// func AtomicAnd32Release(dst *int32, val int32)
TEXT ·AtomicAnd32Release(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ANDL BX, (AX)
    RET
// func AtomicAnd32(dst *int32, val int32)
TEXT ·AtomicAnd32AcqRel(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ANDL BX, (AX)
    RET
// func AtomicAnd32(dst *int32, val int32)
TEXT ·AtomicAnd32SeqCst(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ANDL BX, (AX)
    RET


// func AtomicCAS32(dst *int32, old *int32, new int32)
TEXT ·AtomicCAS32(SB), NOSPLIT, $0-17
    MOVQ dst+0(FP), BX
    MOVL old+8(FP), AX
    MOVL new+12(FP), DX
    CMPXCHGL DX, (BX)
    SETEQ ret+16(FP)
    RET
    

