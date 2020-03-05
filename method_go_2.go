package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y + v.Y)
}

// method can also be written as
// func Abs(v Vertex) float64{
// return math.Sqrt(v.X*v.X + v.Y + v.Y)
// }

// declare a method on non-struct type
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	d := Vertex{3, 4}
	fmt.Println(d.Abs())
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
