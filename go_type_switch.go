package main

import (
	"fmt"
)

func main() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break
		// the break blocks the next statement
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}
}
