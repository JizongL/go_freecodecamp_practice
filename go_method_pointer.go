package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//  if not passed as pointer, the method can't change the values of Vertex struct.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	p := &Vertex{3, 4}
	// p := &v // this will point to v instead of Vertex{3,4}
	(*p).Scale(20)
	fmt.Println(v.Abs())
	fmt.Println(p.Abs())
}
