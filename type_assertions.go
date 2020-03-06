package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	f, ok := i.(float64)
	fmt.Println(f, ok)
	// without ok, it will create panic.
	// f = i.(float64)
	// fmt.Println(f)
	do(3)
	do("string do")
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v %T\n", v, v)
	case string:
		fmt.Printf("%v %T\n", v, v)
	}
}
