package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)

	g.greet2()
	fmt.Println("The new name is:", g.name)
}

type greeter struct {
	greeting string
	name     string
}

// binding this function to the greeter struct
// similar to js prototype binding

// value receiver
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	// g.name won't effect out scope
	// the print on line 13 will still print the originl name, not the empty string.
	g.name = ""
}

// this pass greeter as pointer
func (g *greeter) greet2() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
