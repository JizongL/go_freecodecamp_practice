package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	// delay the function, defer keyword is building a stack, so
	// last one in will be first one go out, assume if all statements
	// are set up with defer, then the last one will be excuted first.
	defer fmt.Println("middle")
	fmt.Println("end")
}
