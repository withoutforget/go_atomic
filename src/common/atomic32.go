package common

func AtomicStore32(dst *int32, src int32)
func AtomicLoad32(src *int32) int32
func AtomicAdd32(dst *int32, val int32)
func AtomicOr32(dst *int32, val int32)
func AtomicAnd32(dst *int32, val int32)
func AtomicCAS32(dst *int32, old *int32, new int32) bool
