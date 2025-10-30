package common

import _ "unsafe"

//go:noescape
func AtomicStore32(dst *int32, src int32)

func AtomicLoad32(src *int32) int32

func AtomicAdd32(dst *int32, val int32)

//go:noescape
func AtomicOr32(dst *int32, val int32)

//go:noescape
func AtomicAnd32(dst *int32, val int32)

//go:noescape
func AtomicCAS32(dst *int32, old *int32, new int32)
