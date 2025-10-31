#include "textflag.h"

// func AtomicStore32(dst *int32, src int32)
TEXT ·AtomicStore32(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    XCHGL BX, (AX)
    RET

// func AtomicLoad32(src *int32) int32
TEXT ·AtomicLoad32(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL (AX), AX
    XCHGL AX, ret+8(FP)
    RET

// func AtomicAdd32(dst *int32, val int32)
TEXT ·AtomicAdd32(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    XADDL BX, (AX)
    RET

// func AtomicOr32(dst *int32, val int32)
TEXT ·AtomicOr32(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    ORL BX, (AX)
    RET

// func AtomicAnd32(dst *int32, val int32)
TEXT ·AtomicAnd32(SB), NOSPLIT, $0-12
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
    

