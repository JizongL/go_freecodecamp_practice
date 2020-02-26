package main

import (
	"fmt"
)

// package level
const c int16 = 43

// iota is a counter
const (
	tt = iota
	bb = iota
	cc = iota
)

// iota can detect pattern
const (
	gg = iota
	mm
	nn
)

func main() {
	// function level

	// type constant
	const myConst int = 42
	const a string = "str"
	const b float32 = 3.14
	const c int = 4

	fmt.Printf("%v,%T\n", myConst, myConst)
	// illegal
	// const myConst float64 = math.Sin(1.57)
	// constant is set at compile time
	// math.Sin() needs to happen during runtime.

	// inner scope shadow outter scope, see const c.

	// compiler can infer type
	const g = 32
	var o int16 = 40
	fmt.Printf("%v,%T\n", g+o, g+o) // result will be inferred to be int16 type

	// inumerate constant
	fmt.Printf("%v,%T\n", tt, tt)
	fmt.Printf("%v,%T\n", bb, bb)
	fmt.Printf("%v,%T\n", cc, cc)
	fmt.Printf("%v,%T\n", tt, tt)
	// inumerate constant
	fmt.Printf("%v,%T\n", gg, gg)
	fmt.Printf("%v,%T\n", mm, mm)
	fmt.Printf("%v,%T\n", nn, nn)

}
