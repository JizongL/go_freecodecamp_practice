package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	// this will print nil
	fmt.Println(ms)
	ms = new(myStruct)
	// deferencing * has lower precedence than ., so we need ()
	// (*ms).foo = 42
	// because go compiler can recognize that ms.foo is not accessing
	// foo from pointer, but it's actually accessing myStruct, so no need to
	// use * in this cases
	ms.foo = 42
	fmt.Println(ms)

	// difference between array and slice with pointers
	arr := [3]int{1, 2, 3}
	b := arr
	fmt.Println(arr, b)
	// this change only arr, not b, because is a brand new copy
	arr[1] = 42
	fmt.Println(arr, b) //[1 42 3] [1 2 3]

	// but this behaves differently with slice.
	sliceArr := []int{1, 2, 3}
	bSlice := sliceArr
	fmt.Println(sliceArr, bSlice)
	sliceArr[1] = 42
	fmt.Println(sliceArr, bSlice) //[1 42 3] [1 42 3]
}

type myStruct struct {
	foo int
}
