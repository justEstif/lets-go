package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func basicTypes() {
	fmt.Printf("Type: %T value %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value %v\n", z, z)
}
