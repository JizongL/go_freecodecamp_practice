package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	// embedding Animal into Bird, this is a composition, not inheritance
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 44
	b.CanFly = false
	fmt.Println(b)
}
