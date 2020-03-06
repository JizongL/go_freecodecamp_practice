package main

import (
	"fmt"
)

func main() {
	var i interface{}
	i = 42
	describe(i)
	i = "string"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v,%T)\n", i, i)
}
