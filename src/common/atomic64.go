package common

func AtomicStore64(dst *int64, src int64)
func AtomicLoad64(src *int64) int64
func AtomicAdd64(dst *int64, val int64)
func AtomicOr64(dst *int64, val int64)
func AtomicAnd64(dst *int64, val int64)
func AtomicCAS64(dst *int64, old *int64, new int64) bool
