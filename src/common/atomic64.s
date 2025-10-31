#include "textflag.h"

// func AtomicStore64(dst *int64, src int64)
TEXT ·AtomicStore64(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    MOVQ BX, (AX)
    RET

// func AtomicLoad64(src *int64) int64
TEXT ·AtomicLoad64(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ (AX), AX
    MOVQ AX, ret+8(FP)
    RET

// func AtomicAdd64(dst *int64, val int64)
TEXT ·AtomicAdd16(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    LOCK
    XADDQ BX, (AX)
    RET

// func AtomicOr64(dst *int64, val int64)
TEXT ·AtomicOr64(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    LOCK
    ORQ BX, (AX)
    RET

// func AtomicAnd64(dst *int64, val int64)
TEXT ·AtomicAnd64(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    LOCK
    ANDQ BX, (AX)
    RET

// func AtomicCAS64(dst *int64, old *int64, new int64) bool
TEXT ·AtomicCAS64(SB), NOSPLIT, $0-25
    MOVQ dst+0(FP), BX
    MOVQ old+8(FP), AX
    MOVQ new+16(FP), DX
    MOVQ (AX),AX
    CMPXCHGQ DX, (BX)
    SETEQ ret+24(FP)
    RET
    

