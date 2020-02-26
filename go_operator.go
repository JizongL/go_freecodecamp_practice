package main

import (
	"fmt"
)

func main() {
	// every time a variable is initialized, it's false by default
	var n bool
	fmt.Printf("%v %T\n", n, n)
	fmt.Println(5 % 4)
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	c := 8 // 2^3
	fmt.Println(c << 3)
	fmt.Println(c >> 3)

	//  floating point
	m := 3.14
	// n = 13.7e72
	// n=2.1E14
	fmt.Printf("%v, %T\n", m, m)

	// complex number
	var t complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", real(t), real(t))
	fmt.Printf("%v, %T\n", imag(t), imag(t))

	var p complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", p, p)

	// string
	s := "this is a string"
	fmt.Printf("%v %T\n", s[2], s[2])

	// string concatenation
	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)
	j := []byte(s2)
	fmt.Printf("%v,%T\n", j, j)

	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}
