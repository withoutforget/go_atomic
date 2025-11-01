#include "textflag.h"

// func AtomicStore64Relaxed(dst *int64, src int64)
TEXT ·AtomicStore64Relaxed(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    MOVQ BX, (AX)
    RET

// func AtomicStore64Release(dst *int64, src int64)
TEXT ·AtomicStore64Release(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    MOVQ BX, (AX)
    RET

// func AtomicStore64SeqCst(dst *int64, src int64)
TEXT ·AtomicStore64SeqCst(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    XCHGQ BX, (AX)
    RET


// func AtomicLoad64Relaxed(src *int64) int64
TEXT ·AtomicLoad64Relaxed(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ (AX), AX
    MOVQ AX, ret+8(FP)
    RET

// func AtomicLoad64Acquire(src *int64) int64
TEXT ·AtomicLoad64Acquire(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ (AX), AX
    MOVQ AX, ret+8(FP)
    RET

// func AtomicLoad64SeqCst(src *int64) int64
TEXT ·AtomicLoad64SeqCst(SB), NOSPLIT, $0-16
    MOVQ x+0(FP), AX
    MOVQ (AX), AX
    XCHGQ AX, ret+8(FP)
    RET


// func AtomicAdd64Relaxed(dst *int64, val int64) int64
TEXT ·AtomicAdd64Relaxed(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    LOCK
    XADDQ BX, (AX)
    MOVQ BX, ret+16(FP)
    RET

// func AtomicAdd64Acquire(dst *int64, val int64) int64
TEXT ·AtomicAdd64Acquire(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    LOCK
    XADDQ BX, (AX)
    MOVQ BX, ret+16(FP)
    RET

// func AtomicAdd64Release(dst *int64, val int64) int64
TEXT ·AtomicAdd64Release(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    LOCK
    XADDQ BX, (AX)
    MOVQ BX, ret+16(FP)
    RET

// func AtomicAdd64AcqRel(dst *int64, val int64) int64
TEXT ·AtomicAdd64AcqRel(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    LOCK
    XADDQ BX, (AX)
    MOVQ BX, ret+16(FP)
    RET

// func AtomicAdd64SeqCst(dst *int64, val int64) int64
TEXT ·AtomicAdd64SeqCst(SB), NOSPLIT, $0-24
    MOVQ x+0(FP), AX
    MOVQ y+8(FP), BX
    LOCK
    XADDQ BX, (AX)
    MOVQ BX, ret+16(FP)
    RET

// func AtomicOr64Relaxed(dst *int64, val int64) int64
TEXT ·AtomicOr64Relaxed(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ORQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET

// func AtomicOr64Acquire(dst *int64, val int64) int64
TEXT ·AtomicOr64Acquire(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ORQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET


// func AtomicOr64Release(dst *int64, val int64) int64
TEXT ·AtomicOr64Release(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ORQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET


// func AtomicOr64AcqRel(dst *int64, val int64) int64
TEXT ·AtomicOr64AcqRel(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ORQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET


// func AtomicOr64SeqCst(dst *int64, val int64) int64
TEXT ·AtomicOr64SeqCst(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ORQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET


// func AtomicAnd64Relaxed(dst *int64, val int64) int64
TEXT ·AtomicAnd64Relaxed(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ANDQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET


// func AtomicAnd64Acquire(dst *int64, val int64) int64
TEXT ·AtomicAnd64Acquire(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ANDQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET


// func AtomicAnd64Release(dst *int64, val int64) int64
TEXT ·AtomicAnd64Release(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ANDQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET


// func AtomicAnd64AcqRel(dst *int64, val int64) int64
TEXT ·AtomicAnd64AcqRel(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ANDQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET


// func AtomicAnd64SeqCst(dst *int64, val int64) int64
TEXT ·AtomicAnd64SeqCst(SB), NOSPLIT, $0-24
    MOVQ dst+0(FP), BX
    MOVQ val+8(FP), CX
loop:
    MOVQ (BX), AX
    MOVQ AX, DX
    ANDQ CX, DX
    CMPXCHGQ DX, (BX)
    JNE loop
    MOVQ AX, ret+16(FP)
    RET



// func AtomicCAS64(dst *int64, old int64, new int64) bool
TEXT ·AtomicCAS64(SB), NOSPLIT, $0-25
    MOVQ dst+0(FP), BX
    MOVQ old+8(FP), AX
    MOVQ new+16(FP), DX
    LOCK
    CMPXCHGQ DX, (BX)
    SETEQ ret+24(FP)
    RET
    
