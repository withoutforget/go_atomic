package common

func AtomicStore64Relaxed(dst *int64, src int64)
func AtomicStore64Release(dst *int64, src int64)
func AtomicStore64SeqCst(dst *int64, src int64)

func AtomicLoad64Relaxed(src *int64) int64
func AtomicLoad64Acquire(src *int64) int64
func AtomicLoad64SeqCst(src *int64) int64

func AtomicAdd64Relaxed(dst *int64, val int64)
func AtomicAdd64Acquire(dst *int64, val int64)
func AtomicAdd64Release(dst *int64, val int64)
func AtomicAdd64AcqRel(dst *int64, val int64)
func AtomicAdd64SeqCst(dst *int64, val int64)

func AtomicOr64Relaxed(dst *int64, val int64)
func AtomicOr64Acquire(dst *int64, val int64)
func AtomicOr64Release(dst *int64, val int64)
func AtomicOr64AcqRel(dst *int64, val int64)
func AtomicOr64SeqCst(dst *int64, val int64)

func AtomicAnd64Relaxed(dst *int64, val int64)
func AtomicAnd64Acquire(dst *int64, val int64)
func AtomicAnd64Release(dst *int64, val int64)
func AtomicAnd64AcqRel(dst *int64, val int64)
func AtomicAnd64SeqCst(dst *int64, val int64)

func AtomicCAS64(dst *int64, old int64, new int64) bool
