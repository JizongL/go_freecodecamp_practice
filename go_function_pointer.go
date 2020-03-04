package main

import (
	"fmt"
)

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// variadic parameter, similar to spread operator.
func sum(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
	// return a pointer value
	// heap memory, it won't be cleared.
	// return &result
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		// return two values, 0 and error
		return 0.0, fmt.Errorf("Can not divide by zero")
	}
	return a / b, nil
}

func main() {
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
	// common error handler
	d, err := divide(5.0, 0.0)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
