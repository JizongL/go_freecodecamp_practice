package main

import (
	"fmt"
)

func main() {
	// array can only hold one type of element

	// grades := [3]int{97, 85, 93}
	// grades2 := [...]int{34, 32, 34, 44}

	// empty array
	var students [3]string
	// assign array value
	students[0] = "string 1"

	fmt.Printf("Grades: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students))
}
