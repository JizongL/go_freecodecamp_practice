package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		// if the func is dealing with async code inside, then
		// i is passed in and thus has inner scope, so it won't
		// be effected by the outter scope i as the loop is moving
		// forward.
		func(i int) {
			fmt.Println(i)
		}(i)
	}
}
