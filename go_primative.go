package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v,%T\n", j, j)

	// convert int to unitcode, so 42 is an asterisk
	var k string
	k = string(i)
	fmt.Printf("%v %T\n", k, k)

	var l string
	l = strconv.Itoa(i)
	fmt.Printf("%v %T\n", l, l)

	// 3 ways to declare variables
	var foo int
	var bar int = 42
	doo := 43

	// can't redeclare variables, but can shadoew them
	//  all variables must be used

	// visibility
	//  lower case first letter for package scope
	//  upper case first letter to export
	//  no private scope

	// go doesn't do implicit conversion.
}
