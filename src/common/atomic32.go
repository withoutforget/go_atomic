package common

func AtomicStore32Relaxed(dst *int32, src int32)
func AtomicStore32Release(dst *int32, src int32)
func AtomicStore32SeqCst(dst *int32, src int32)

func AtomicLoad32Relaxed(src *int32) int32
func AtomicLoad32Acquire(src *int32) int32
func AtomicLoad32SeqCst(src *int32) int32

func AtomicAdd32Relaxed(dst *int32, val int32) int32
func AtomicAdd32Acquire(dst *int32, val int32) int32
func AtomicAdd32Release(dst *int32, val int32) int32
func AtomicAdd32AcqRel(dst *int32, val int32) int32
func AtomicAdd32SeqCst(dst *int32, val int32) int32

func AtomicOr32Relaxed(dst *int32, val int32) int32
func AtomicOr32Acquire(dst *int32, val int32) int32
func AtomicOr32Release(dst *int32, val int32) int32
func AtomicOr32AcqRel(dst *int32, val int32) int32
func AtomicOr32SeqCst(dst *int32, val int32) int32

func AtomicAnd32Relaxed(dst *int32, val int32) int32
func AtomicAnd32Acquire(dst *int32, val int32) int32
func AtomicAnd32Release(dst *int32, val int32) int32
func AtomicAnd32AcqRel(dst *int32, val int32) int32
func AtomicAnd32SeqCst(dst *int32, val int32) int32

func AtomicCAS32(dst *int32, old *int32, new int32) bool
