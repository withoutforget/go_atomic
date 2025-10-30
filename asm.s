#include "textflag.h" // Includes some standard macros.
// func AtomicStore32(dst*Atomic32, src Atomic32)
TEXT ·AtomicStore32(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    MOVL BX, (AX)
    RET

// func AtomicLoad32(src *Atomic32) Atomic32
TEXT ·AtomicLoad32(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL (AX), AX
    MOVL AX, ret+8(FP)
    RET

// func AtomicAdd32(dst *Atomic32, val int32)
TEXT ·AtomicAdd32(SB), NOSPLIT, $0-12
    MOVQ x+0(FP), AX
    MOVL y+8(FP), BX
    LOCK
    XADDL BX, (AX)
    RET

