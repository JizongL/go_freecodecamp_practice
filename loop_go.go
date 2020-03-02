package main

import (
	"fmt"
)

func main() {
	// where there is only one variable, you can do i++, but not when there are two or more
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i)
		fmt.Println(j)
	}
	// or can define i first and leave the first place in the for expression empty
	i := 2
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	// or move the i++ inside the block, but don't remove the ; in the for expression
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// go has no while loop, but can use for this way
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
	// with continues keyword
	// when the following is run, all the above should be disable.
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	// use a Loop to designate which loop to break
Loop:
	for i := 2; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				// break the entire loop, not just the inner one.
				break Loop
			}
		}
	}

	// for range loop
	s := []int{1, 2, 3}
	for k, v := range s {
		// k is position, v is value
		fmt.Println(k, v)
	}

	// loop through string
	// without string conversion, it will print out number
	// key has to be present, if not wanted, use _ to serve as place holder
	str := "hello Go!"
	for k, v := range str {
		fmt.Println(k, string(v))
	}
}
