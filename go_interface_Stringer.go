package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	p1 := Person{"Tom", 23}
	p2 := Person{"John", 32}
	fmt.Println(p1, p2)

}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v year)", p.Name, p.Age)
}
