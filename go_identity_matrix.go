package main

import (
	"fmt"
)

func main() {
	// array is considered value

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)
	// array is considered value
	arr := [...]int{1, 2, 3}
	arrTwo := arr
	arrTwo[1] = 4
	fmt.Println(arr, arrTwo)
	// use pointer, don't just reference array
	arrThree := &arr
	arrThree[1] = 8
	fmt.Println(arr, arrThree)
	// slice, is a reference type
	sliceArr := []int{1, 2, 3}
	sliceArrCopy := sliceArr
	sliceArrCopy[1] = 9 // reference
	fmt.Printf("Capacity: %v\n", cap(sliceArr))
	fmt.Println(sliceArr, sliceArrCopy) // same result

	// slicing
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// a := [...]int{1, 2, 3, 4, 5, 6, 7, 8} change slice to array with ...
	// slice works for both array and slice.
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	// if the 3rd argument 100 is not passed, the cap will be default 3.
	makeArr := make([]int, 3, 100)
	fmt.Println(makeArr)
	fmt.Printf("Length %v\n", len(makeArr))
	fmt.Printf("Capacity: %v\n", cap(makeArr))

	// append
	appendArr := []int{}
	fmt.Println(appendArr)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	appendArr = append(appendArr, 1)
	fmt.Println(cap(appendArr))

	// ... act as if Javascript spread
	// a = append(a,[]int{2,3,4,5}...)
	// b:=a[:len(a)-1]
	// if want to remove element in between
	// 1, 2, 3, 4, 5, 6, 7, 8
	// 3 is removed
	bb := append(a[:2], a[3:]...)
	fmt.Println(bb)

}
